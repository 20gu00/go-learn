package main

//引入orm包时，需要同时引入驱动包
//"gorm.io/driver/mysql"就是gorm，mysql的驱动包
import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//表对应的struct
type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
	Password string
}

//自定义表名
func (*User) TableName() string {
	//返回表名
	return "user_t"
}

//数据库初始化
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动创建表,AutoMigrate没有建立映射关系
	//用于项目初始化自动创建表
	//db.AutoMigrate(User{})

	//使用结构体的时候会建立映射关系,如果有自定义表名就去找该表

	//新增
	//使用的是指针否则会报错
	//data := User{
	//  id自增
	//	Username: "zhangsan2",
	//	Password: "123456",
	//}
	//接收的全都是指针类型的interface
	//db.Create(&data)

	//修改
	//var id = 1
	//where 里面是原生的查询条件
	//?代表变量，可将值传入
	//没有使用结构体就要指定表名，通常使用Model(User{}),或Table("user_t")
	//db.Model(User{}).Where("id = ?", id).Update("username", "lisi")
	//db.Model(User{ID: 1}).Update("password", "666")
	//db.Table("user_t").Where("id = ?", id).Update("password", "666")

	//查询
	//第一种
	u := &User{ID: 3}
	db.First(u)
	fmt.Println(u)
	//第二种用法
	u2 := &User{}
	//db.Where("id = ?", 3).First(u2)
	fmt.Println(u2)
	//查询所有数据
	users := []User{}
	db.Find(&users)
	fmt.Println(users)
	//删除,delete方法中传入的是结构体
	//注意删除整张表
	//db.Where("id = ?", 3).Delete(&User{})
	//错误处理
	tx := db.Where("id = ?", 1).First(u2)
	//注意一点，tx.Error中，为空也会被算成error  "record not found"
	//tx.Error==gorm.ErrRecordNotFound
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println(tx.Error)
	}
}
