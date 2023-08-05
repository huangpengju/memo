package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(Db, DSN string) {
	db, err := gorm.Open(Db, DSN)
	if err != nil {
		fmt.Printf("数据库连接错误：= %v", err)
		panic("数据库连接错误")
	}
	fmt.Println("+++++++++数据库连接成功+++++++++++")
	// 开启打印日志 （打印 sql 功能）
	db.LogMode(true)

	// 判断 gin 框架是不是发行版
	if gin.Mode() == "release" {
		// 是发行版，不打印日志 SQL
		db.LogMode(false)
	}
	// 建表的时候表名字不加s
	db.SingularTable(true)
	// 设置连接池
	db.DB().SetMaxIdleConns(20)
	// 最大连接数
	db.DB().SetMaxOpenConns(100)
	// 连接时间
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
}
