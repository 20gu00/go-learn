package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //记得导入驱动,匿名引用,本包不调用这个名,但import是这个包执行了常量变量init
)

type User struct {
	id   int
	name string
	age  int
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

	//查询,记得释放数据库连接

	//单行查询
	sqlStr := "select id,name,age from user where id=?"
	var u User
	//如果没有scan,持有的数据库连接不会被释放
	//如果恰巧连接书被限制了很可能导致新的需要获取连接的操作hang住
	//if err:=db.QueryRow(sqlStr,1).Scan(&u.id,&u.name,&u.age);err!=nil{
	row := db.QueryRow(sqlStr, 1)
	if err := row.Scan(&u.id, &u.name, &u.age); err != nil {
		fmt.Printf("scan failed err:%v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)

	//多行查询
	sqlStr = "select id,name,age from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed err:%v", err)
		return
	}

	//关闭rows,释放持有的数据库连接
	defer rows.Close()

	for rows.Next() { //每一行记录
		if err := rows.Scan(&u.id, &u.name, &u.age); err != nil {
			fmt.Printf("scan failed err:%v", err)
		}

		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}

	//新增更新删除
	sqlStr = "insert into user(name,age) values (?,?)"
	ret, err := db.Exec(sqlStr, "aa", 100) //结果
	if err != nil {
		fmt.Printf("插入数据失败 err:%v\n", err)
		return
	}

	//最后一次插入数据的id
	theId, err := ret.LastInsertId() //插入数据的id
	if err != nil {
		fmt.Printf("获取数据的id失败 err:%v\n", err)

	}
	fmt.Printf("插入数据成功 id:%v\n", theId)

	//更新数据
	sqlStr = "update user set age=? where id=?"
	ret, err = db.Exec(sqlStr, 1000, 1) //结果
	if err != nil {
		fmt.Printf("更新数据失败 err:%v\n", err)
		return
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取更新的行数失败")
		return
	}

	fmt.Println("影响的行数", n)

	//删除
	sqlStr = "delete from user where id=?"
	ret, err = db.Exec(sqlStr, 3) //结果
	if err != nil {
		fmt.Printf("更新数据失败 err:%v\n", err)
		return
	}

	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Println("获取更新的行数失败")
		return
	}

	fmt.Println("删除成功,影响的行数", n)
}
