package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// SaveImgCode 存储图片id到redis
func SaveImgCode(code, uuid string) error {
	// 连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn err:", err)
		return err
	}
	defer conn.Close()
	// 操作数据库
	_, err = conn.Do("setex", uuid, 60*5, code)
	return err
}
