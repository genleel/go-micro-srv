package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// 设置临时session
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/cookie", func(context *gin.Context) {
		// cookie设置
		context.SetCookie("mytest", "hello", 60*60, "", "", true, true)
		context.Writer.WriteString("测试cookie.")
		// cookie获取
		cookie, _ := context.Cookie("mytest")
		fmt.Println(cookie)
	})

	router.GET("/session", func(context *gin.Context) {
		// 调用session设置数据
		session := sessions.Default(context)

		// 设置session
		session.Set("hello", "world")

		// 获取session
		v := session.Get("hello")
		fmt.Println("获取session, ", v)

		// 修改session需要save否则不生效
		session.Save()
		context.Writer.WriteString("测试session.")
	})

	router.Run(":9999")
}
