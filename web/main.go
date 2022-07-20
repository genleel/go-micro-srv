package main

import (
	"github.com/gin-gonic/gin"
	"go-micro-srv/web/controller"
	"go-micro-srv/web/model"
)

func main() {

	// 初始化 Redis 链接池
	model.InitRedis()

	router := gin.Default()

	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目启动...")
	})*/

	router.Static("/home", "view")
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
	}

	router.Run(":8080")
}
