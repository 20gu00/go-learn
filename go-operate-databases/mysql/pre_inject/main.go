package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //unknown driver "mysql" (forgotten import?)
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/gateway"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("建立mysql连接中")

	if err := db.Ping(); err != nil {
		fmt.Println("连接数据库失败", err) //%v
		panic(err)
	}
	fmt.Println("建立mysql连接成功")

	//prepare
	//*生产环境慎
	sqlStr := "select * from user where id>?"
	//可用的状态
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理失败", err.Error())
		return
	}

	defer stmt.Close() //如果stmt为nil,会出错
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("查询失 err:%v\n", err)
		return
	}

	//var u:= *User //使用指针 nil 未初始化,没有分配内存
	//u:=&User{}  make:chan slice map引用类型的数据
	u := new(User) //指针
	for rows.Next() {
		//成员也一样指针
		//顺序
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			fmt.Printf("遍历记录失败 err:%v\n", err)
			return
		}
		fmt.Printf("id:%v age:%v name%v\n", u.Id, u.Age, u.Name)
	}
}
