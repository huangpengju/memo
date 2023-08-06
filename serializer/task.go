package serializer

import "memo/model"

// Taks 是序列化 json 的专用结构
type Task struct {
	ID        uint   `json:"id" example:"1"` // 备忘录id
	Title     string `json:"title" example:"学习"`
	Content   string `json:"content" example:"Go编程"`
	Status    int    `json:"status" example:"0"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}
