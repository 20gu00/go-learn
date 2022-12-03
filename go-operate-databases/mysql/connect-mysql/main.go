package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //记得导入驱动,匿名引用,本包不调用这个名,但import是这个包执行了常量变量init
)

func main() {
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/gateway"
	//返回的DB对象可以被多个goroutine安全并发使用,自己维护了连接池,Open一般只调用一次,DB对象也很少需要被关闭
	db, err := sql.Open("mysql", dsn) //并不是真正连接了数据库,只是校验参数格式是否正确
	if err != nil {
		panic(err)
	}
	defer db.Close() //确保db是正常的,释放数据库连接相关的资源
	fmt.Println("建立mysql连接中")

	//初始化数据库连接,真正建立数据库连接
	if err := db.Ping(); err != nil { //尝试建立数据库连接,校验dsn是否正确
		fmt.Println("连接数据库失败", err) //%v
		panic(err)
	}

	fmt.Println("建立mysql连接成功")

	//db:
	//db.SetMaxOpenConns(n)  设置go程序和数据库见能建立多少连接数,默认0(n<=0)不限制
	//db.SetMaxIdleConns(n)  同
	//db.SetMaxLifeTime(1*time.Second)  连接最长的时间
}
