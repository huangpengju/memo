// model 包是应用数据库模型
// user.go 用于用户的创建
package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

// SetPassword 加密密码
// 方法的接收者 User 模型
// 方法的参数 password 是未加密的密码
// 方法的返回值 error 错误
func (user *User) SetPassword(password string) error {
	// bcrypt 包
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// CheckPassword 对比密码
// 方法的接收者 User 模型
// 方法的参数 password 是未加密的密码
// 方法的返回值 true（密码相等） 或 false(密码不相等)
func (user *User) CheckPassword(password string) bool {
	// CompareHashAndPassword比较bcrypt哈希后的密码与可能的密码明文等价。成功时返回nil，失败时返回错误。
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
