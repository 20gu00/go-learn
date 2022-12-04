package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func main() {
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&&parseTime=True&&loc=Local" //Asia/Shanghai数据库时区设置

	//建立连接
	//MustConnect连接不成功就直接panic,普通的连接会返回错误开发者决定处理
	DB, err1 := sqlx.Connect("mysql", dsn) //原生的Open+Ping
	if err1 != nil {
		fmt.Printf("连接数据库失败 err:%v\n", err1)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	fmt.Println("数据库连接成功")

	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("开启事务", err)
		return
	}

	//在这里捕获panic和rollback
	//当前函数退出之前
	defer func() {
		if p := recover(); p != nil { //recover()要在defer
			fmt.Println("rollback")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				fmt.Println("rollback", err)
				tx.Rollback() //依旧可以rollback
			}
			fmt.Println("commit")
		}
	}()

	sqlStr := "update user set age=10000 where id = ?"
	ret, err := DB.Exec(sqlStr, 3) //这个err只是判断Exec执行是不是失败,如果没有这条记录也就是sql的错误实在ret中
	if err != nil {
		return
	}

	//判断这次执行影响的记录,成功的话就是影像一条,不然就是0
	n, err := ret.RowsAffected() //这个err只是判断这里是否执行成功
	if err != nil {
		return
	}

	if n != 1 {
		fmt.Println("affect not 1")
		fmt.Println(n)
		//tx.Rollback()
		err = errors.New("n not 1") //设置err
		return
	}

	ret, err = DB.Exec(sqlStr, 10)
	if err != nil {
		return
	}

	n, err = ret.RowsAffected()
	if err != nil {
		return
	}

	if n != 1 {
		fmt.Println("affect not 1")
		fmt.Println(n)
		err = errors.New("n not 1")
		return
	}

	//err最好区分
	//顺序处理rollback更易读
}
