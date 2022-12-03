package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB //并发安全,直接共用即可(大写,跨包调用)

func InitDb() error {
	dsn := "root:100.Acjq@tcp(127.0.0.1:6379)/gateway"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer db.Close()
	fmt.Println("建立mysql连接中")
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
