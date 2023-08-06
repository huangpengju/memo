package serive

import (
	"fmt"
	"memo/model"
	"memo/serializer"
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
