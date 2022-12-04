package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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

	//Tx是一个正在进行的数据库事务。事务必须以调用Commit或Rollback结束。
	tx, err := db.Begin()
	if err != nil {
		if tx != nil { //tx!=nil才能操作,不然会报错
			tx.Rollback()
		}
		fmt.Println(err)
		return
	}

	sqlStr := "update user set age=111 where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理失败", err.Error())
		return
	}

	defer stmt.Close()
	ret1, err := stmt.Exec(1)
	row1, _ := ret1.RowsAffected()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return //可以return
	}

	ret2, err := stmt.Exec(2)
	row2, _ := ret2.RowsAffected()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return //可以return
	}

	if row1 == 1 && row2 == 1 { //其实不需要,上面已经判断了错误,有错误就会滚并且return了
		if err = tx.Commit(); err != nil {
			tx.Rollback() //提交失败也回滚
			fmt.Println(err)
			return
		}
	} else {
		tx.Rollback()
		fmt.Println("事务回滚")
		return
	}

	fmt.Println("事务完成")
}
