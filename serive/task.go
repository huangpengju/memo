package service

import (
	"memo/model"
	"memo/serializer"
	"time"
)

// CreateTaskService 结构体用于创建备忘录时用于接收请求中的参数。
type CreateTaskService struct {
	Title   string `from:"title" json:"title" form:"title"`
	Content string `from:"content" json:"content" form:"content"`
	Status  int    `from:"status" json:"status" form:"status"` // 0是未完成, 1是已完成
}

// Create
func (service *CreateTaskService) Create(id uint) serializer.Response {
	// 用户的结构
	var user model.User
	// 查询用户的信息
	model.DB.First(&user, id)
	code := 200
	// 准备 备忘录结构各字段的数据
	task := model.Task{
		Uid:       user.ID,
		Title:     service.Title,
		Status:    service.Status,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500 //失败
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建备忘录成功",
	}
}
