// model 包是应用数据库模型
// user.go 用于用户的创建
package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}
