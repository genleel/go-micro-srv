package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建全局结构体
type Student struct {
	gorm.Model
	Id   int // 默认主键
	Name string
	Age  int
}

// 创建全局连接池
var GlobalConn *gorm.DB

func main() {
	// 连接数据库
	dsn := "ligen:LiGen1129!@tcp(127.0.0.1:3306)/go_micro_srv?charset=utf8mb4&parseTime=True&loc=Local"
	GlobalConn, _ = gorm.Open(mysql.Open(dsn))

	db, _ := GlobalConn.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	// gorm创建表
	/*err := GlobalConn.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println("Migrate err: ", err)
		return
	}*/

	InsertData()

}

func InsertData() {
	// 创建数据
	var stu Student
	stu.Name = "张三"
	stu.Age = 20

	// 插入数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {

	// GlobalConn.First(&stu)   查询 第一条全部信息.
	// GlobalConn.Select("name, age").First(&stu)   查询第一条 name 和 age
	// GlobalConn.Select("name, age").Find(&stu)   查询所有条记录的  name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)  查询姓名为lisi的 name 和 age
	// GlobalConn.Select("name, age").Where("name = ?", "lisi").Where("age = ?", 22).Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ? and age = ?", "lisi", 22).Find(&stu)
	//GlobalConn.Where("name = ?", "lisi").Select("name, age").Find(&stu)
	//GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)

	var stu []Student
	GlobalConn.Unscoped().Find(&stu)
	fmt.Println(stu)
}

func UpdateData() {
	/*	var stu Student

		stu.Name = "wangwu"
		stu.Age = 77
		stu.Id = 4*/

	/*	fmt.Println(GlobalConn.Model(new(Student)).Where("name = ?", "zhaoliu").
		Update("name", "lisi").Error)*/

	fmt.Println(GlobalConn.Model(new(Student)).Where("age = ?", 77).
		Updates(map[string]interface{}{"name": "lisi", "age": 119}).Error)
}

func DeleteData() {
	fmt.Println(GlobalConn.Unscoped().Where("name = ?", "lisi").Delete(new(Student)).Error)
}
