package route

import (
	"fishba.top/controller"
	"fishba.top/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine)*gin.Engine {
	// 加载html文件，即template包下所有文件
	r.Static("/static","static")
	r.LoadHTMLGlob("template/*")

	//加载用户认证中间件

	//用户管理相关路由
	user:=r.Group("/api/auth")
	{
		user.POST("/register",controller.Register)
		user.POST("/login",controller.Login)
		user.GET("/info",middleware.AuthMiddleware(),controller.Info)
	}


	//运维LOG相关路由
	r.Use(middleware.AuthMiddleware())
	job:=r.Group("/job")
	{
		job.GET("/",controller.ShowJob)
		job.POST("/add",controller.AddJob)
		job.GET("/add",controller.AddJob)
		job.PUT("/update",controller.UpdateJob)
		job.DELETE("/del",controller.DelJob)
	}
	return r
}
