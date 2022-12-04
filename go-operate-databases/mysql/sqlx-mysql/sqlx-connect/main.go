package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func main() {
	//time.Time<=>datatime
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&&parseTime=True&&loc=Local" //Asia/Shanghai数据库时区设置

	//建立连接
	//MustConnect连接不成功就直接panic,普通的连接会返回错误开发者决定处理
	DB, err := sqlx.Connect("mysql", dsn) //原生的Open+Ping
	if err != nil {
		fmt.Printf("连接数据库失败 err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	fmt.Println("数据库连接成功")
	return
}
