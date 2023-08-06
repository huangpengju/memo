package api

import (
	"fmt"
	"memo/serive"

	"github.com/gin-gonic/gin"
)

// UserRegister 函数接收注册路由
// 把接收到的信息传给 userService
func UserRegister(c *gin.Context) {
	// 声明 userRegister 模型，接收请求中的 UserName 和 PassWord
	var userRegister serive.UserService
	// 把请求中的数据与 userRegister 模型进行绑定
	if err := c.ShouldBind(&userRegister); err == nil {
		// 绑定 userRegister 数据后，开始注册用户
		res := userRegister.Register()
		fmt.Println("注册绑定成功，请看注册返回===", res)

		c.JSON(200, res)
	} else {
		fmt.Println("注册绑定失败===", err)
		c.JSON(400, err)
	}
}
