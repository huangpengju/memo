package api

import (
	"fmt"
	"memo/pkg/utils"
	service "memo/serive"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// CreateTask 创建 备忘录
func CreateTask(c *gin.Context) {
	// 声明 createTask 结构接受 请求中的参数
	var createTask service.CreateTaskService

	// 验证 token
	// 返回 token 所有者的信息
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	// 把请求中的数据与createTask结构中的字段绑定
	if err := c.ShouldBind(&createTask); err == nil {
		fmt.Println("createTask====", createTask)
		res := createTask.Create(claim.Id)

		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

// ShowTask 查询备忘录详情
func ShowTask(c *gin.Context) {
	// 声明 createTask 结构接受 请求中的参数
	var showTask service.ShowTaskService

	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

// 查询所有备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

// 修改备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService

	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
