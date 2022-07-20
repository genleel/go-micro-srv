package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	defer conn.Close()
	// 操作数据库
	reply, err := conn.Do("set", "hello", "world")

	// 回复助手类函数
	s, err := redis.String(reply, err)
	fmt.Println(s, err)
}
