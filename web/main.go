package main

import (
	"github.com/gin-gonic/gin"
	"go-micro-srv/web/controller"
)

func main() {
	router := gin.Default()

	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目启动...")
	})*/

	router.Static("/home", "view")
	router.GET("/api/v1.0/session", controller.GetSession)
	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd)

	router.Run(":8080")
}
