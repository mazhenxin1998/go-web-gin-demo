package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
定义全局变量.
*/

var db *gorm.DB

type GoTest struct {
	Id   int
	Name string
}

/**
操作数据库进行增加操作.
*/
func Add() {

	// 这个永远是1啊.
	// 连接数据库.
	g1 := new(GoTest)
	g1.Name = "你好,GoLang"
	db.Create(g1)
	fmt.Println("向DB中增加数据成功. ")

}

/**
从DB中查询数据.
*/
func Query(id int) *GoTest {

	result := new(GoTest)
	// db.First(&user, 10)
	////// SELECT * FROM users WHERE id = 10;
	db.First(&result, id)
	// 方法执行到了这里之后,查到了数据. result也就有了数据.
	return result

}

/**
在当前文件被夹在之前进行数据库的连接操作。
*/
func init() {

	d, err := gorm.Open("mysql", "root:mzx123@(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {

		// 有异常.
		fmt.Println("连接数据库出现错误了. ")
		return

	}

	db = d
	fmt.Println("连接数据库成功. ")

}
