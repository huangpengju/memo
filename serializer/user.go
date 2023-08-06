package serializer

import "memo/model"

// User 结构用于登录成功后，响应返回的用户信息
type User struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	UserName  string `json:"user_name" form:"user_name" example:"admin"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
}

// BuildUser 建立用户信息
// 参数是数据库查询到的用户信息，绑定至 user
// 返回值是序列化中定义的 User
func Builduser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
