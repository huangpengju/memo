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
