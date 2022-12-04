package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	//字段类型要对应,大写,sqlx包使用
	//tag 数据库字段名
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
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

	//单行查询
	sqlStr := "select id,age,name from user where id = ?"
	u := new(User)
	//预处理方式
	//stmt,err:=DB.Prepare(sqlStr)
	//注意传递的是结构体指针
	err = DB.Get(u, sqlStr, 1)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}

	fmt.Println(u) //直接打印结构体 指针 *u

	//多行查询
	sqlStr = "select id,age,name from user where id > ?"
	u2 := new([]User)
	if err := DB.Select(u2, sqlStr, 0); err != nil {
		fmt.Println("多行查询失败", err)
		return
	}

	fmt.Println(u2)
	for _, item := range *u2 { //解引用
		fmt.Println(item)
	}

	//新增数据
	sqlStr = "insert into user(name,age) values(?,?)"
	ret, err := DB.Exec(sqlStr, "ww", 100)
	if err != nil {
		fmt.Println("新增数据失败", err)
		return
	}

	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("获取最后一次插入数据的id失败", err)
		return
	}
	fmt.Println("最后一次操作的记录的id", theID)

	//更新数据  条件不然全部更新
	sqlStr = "update user set name=?,age=? where id=?"
	ret, err = DB.Exec(sqlStr, "ww", 100, 1)
	if err != nil {
		fmt.Println("更新数据失败", err)
		return
	}

	theID, err = ret.LastInsertId()
	if err != nil {
		fmt.Println("获取最后一次记录的id失败", err)
		return
	}
	fmt.Println("最后一次记录的id", theID)

	//删除数据
	sqlStr = "delete from user where id=?"
	ret, err = DB.Exec(sqlStr, 2)
	if err != nil {
		fmt.Println("删除数据失败", err)
		return
	}

	theID, err = ret.LastInsertId()
	if err != nil {
		fmt.Println("获取最后一次记录的id失败", err)
		return
	}
	fmt.Println("最后一次记录的id--", theID)

	//namedExec
	_, err = DB.NamedExec("insert into user(name,age) values(:name,:age)",
		map[string]interface{}{
			"name": "aa",
			"age":  100,
		})

	//namedQuery
	//``
	sqlStr = "select id,name,age from user where name=:name"
	rows, err := DB.NamedQuery(sqlStr,
		map[string]interface{}{
			"name": "aa",
		})

	if err != nil {
		fmt.Println("namedQuery failed", err)
		return
	}
	defer rows.Close() //关闭连接
	for rows.Next() {
		u3 := new(User)
		//最好处理scan的错误
		rows.StructScan(u3) //最好指针  MapScan SliceScan
		fmt.Println(u3.ID)
	}

	//第二种写法
	u4 := new(User)
	rows, err = DB.NamedQuery(sqlStr, u4)
	fmt.Println(u4)

	//上面是案例演示,注意不要这样写,最好每个增删改查分装成一个函数来调用,这样避免共用一个连接,共用一个连接时往往查询会有问题,查到的数据都是零值
	//最好每个增删改查分装成一个函数来调用
}
