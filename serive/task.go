package service

import (
	"memo/model"
	"memo/serializer"
	"time"
)

// CreateTaskService 结构体用于创建备忘录时接收请求中的参数。
type CreateTaskService struct {
	Title   string `from:"title" json:"title" form:"title"`
	Content string `from:"content" json:"content" form:"content"`
	Status  int    `from:"status" json:"status" form:"status"` // 0是未完成, 1是已完成
}

// ShowTaskService 结构体用于查询一条备忘录时接受请求中的参数
// ShowTaskService 同时是查询备忘录 Show 方法的接收者
type ShowTaskService struct {
}

// ListTaskService 结构体用于查询所有备忘录时，接收查询条件
// 该结构体拥有 List 方法
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`   // 当前页
	PageSize int `json:"page_size" form:"page_size"` // 每页最多显示多少条
}

// 更新
// Update()方法的接收者
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

// 模糊查询
type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// DeleteTaskService 删除一条备忘录
// DeleteTaskService 该接收者有一个删除的方法
type DeleteTaskService struct {
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

// Show 查询一条备忘录的详细信息
func (service *ShowTaskService) Show(tid string) serializer.Response {
	// 定义个 task 模型，与数据库对接
	var task model.Task

	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库查询这条备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功",
	}
}

// List 列表返回用户所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	// 定义一个切片
	var tasks []model.Task

	count := 0

	// 如果请求中的 PageSize 为0，则默认显示每页10条
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

// 更新备忘录
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	code := 200
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "修改失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "修改成功",
	}
}

// 模糊查询
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	code := 200
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 删除方法
func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task

	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
