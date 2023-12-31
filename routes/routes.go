package routes

import (
	"memo/api"
	"memo/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// NewRouter
// 返回值是 gin 引擎
// 引擎是框架的实例，它包含了复用器、中间件和配置设置。
// 通过New()或Default()创建Engine实例
func NewRouter() *gin.Engine {
	// 创建路由
	r := gin.Default()

	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 路由组1，处理POST请求
	v1 := r.Group("api/v1")
	// {} 是书写规范
	{
		// 用户注册操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		// JWT() 检查 token
		authed.Use(middleware.JWT())
		{
			// 创建备忘录
			authed.POST("task", api.CreateTask)
			// 查询一条备忘录
			authed.GET("task/:id", api.ShowTask)
			// 查询所有的备忘录(POST方法)
			authed.POST("tasks", api.ListTask)
			// 查询所有的备忘录(GET方法)

			// 修改当前备忘录
			authed.PUT("task/:id", api.UpdateTask)

			// 搜索一条备忘录
			authed.POST("search", api.SearchTask)

			// 删除一条备忘录
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return r
}
