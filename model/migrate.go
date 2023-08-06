// model 包是应用数据库模型
// migrate.go 用于模型创建数据库（数据库迁移）
package model

func migration() {
	// 自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})

	// Model()指定要运行数据库操作的模型
	// 设置外键
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
