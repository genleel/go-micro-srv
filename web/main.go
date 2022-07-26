package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-micro-srv/web/controller"
	"go-micro-srv/web/model"
)

func main() {

	// 初始化 Redis 链接池
	model.InitRedis()
	// 初始化 Mysql 连接池
	model.InitDb()

	router := gin.Default()

	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	/*router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目启动...")
	})*/

	router.Use(sessions.Sessions("mysession", store))

	router.Static("/home", "view")
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
		r1.GET("/areas", controller.GetArea)
		r1.POST("/sessions", controller.PostLogin)
		r1.DELETE("/session", controller.DeleteSession)

	}

	router.Run(":8080")
}
