package main

import (
	_ "fishba.top/common"
	"fishba.top/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
func main()  {

	router := gin.Default()
	//相关路由
	route.CollectRoutes(router)
	// 指定地址和端口号
	router.Run(viper.GetString("server.port"))
}

