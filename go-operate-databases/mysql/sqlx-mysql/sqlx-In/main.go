package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//database/sql/driver
func (u *User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

//使用sqlx.In拼接语句和参数
func BatchInsertUser2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"insert into user(name,age) values(?),(?),(?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := DB.Exec(query, args...)
	return err
}
func main() {
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
