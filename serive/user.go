package service

import (
	"fmt"
	"memo/model"
	"memo/pkg/utils"
	"memo/serializer"

	"github.com/jinzhu/gorm"
)

// 用户服务
type UserService struct {
	UserName string `from:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `from:"password" json:"password" binding:"required,min=5,max=16"`
}

// Register 用户注册
// 方法的接收者 UserService
// 返回 Json 格式数据
func (service *UserService) Register() serializer.Response {
	// 声明一个 user 模型
	var user model.User
	var count int

	// Model()指定要运行数据库操作的模型
	// Where First Count 查询是不是有同名的用户存在
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		// 用户名已经存在 返回错误消息
		msg := fmt.Sprintf("已经有“%v”这个人了，无需再注册", service.UserName)
		return serializer.Response{
			Status: 400,
			Msg:    msg,
		}
	}
	// 用户名可用，准备注册

	// 给 User 模型赋值
	user.UserName = service.UserName
	// 对密码加密
	// SetPassword 方法的接收者是 user, 方法内部实现了给 user.PassWordDigest 赋值
	if err := user.SetPassword(service.PassWord); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 上述代码已经给 user 模型中的字段 UserName、PassWordDigest 完成了赋值
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

// Login 用户登录
// 方法的接收者是 UserService
// 方法返回 Json 格式数据
func (service *UserService) Login() serializer.Response {
	var user model.User

	// 先验证用户
	// 查询条件是 接收的参数
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		// 如果记录未找到 err = RecordNotFound
		// 如果error包含RecordNotFound错误 IsRecordNotFoundError 返回 true
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Data:   err,
				Msg:    "用户不存在,请先登录",
			}
		}
		// 用户存在，是其他因素的错误
		return serializer.Response{
			Status: 500,
			Data:   err,
			Msg:    "数据库查询登录用户信息出错",
		}
	}
	// 找到用户后
	// 去验证登录用户的密码
	if !user.CheckPassword(service.PassWord) {
		// 密码不相等时
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	// 密码匹配成功后,响应用户信息和token

	// 准备一个 token ,作为响应返回
	token, err := utils.GenerateToken(user.ID, user.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发出错",
		}
	}
	// 返回用户信息和token
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.Builduser(user), Token: token},
		Msg:    "登录成功",
	}
}
