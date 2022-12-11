package main

//引入orm包时，需要同时引入驱动包
//"gorm.io/driver/mysql"就是gorm，mysql的驱动包
import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 一
type User struct {
	gorm.Model
	Username string
	//或者指针
	CreditCards []*CreditCard //[]*  *[]
}

// 多
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 默认就是使用User表的ID作为这张表的外键进行关联(所有者类型加上主键),类型对应
}

//数据库初始化
func main() {
	//定义连接地址
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(User{}, CreditCard{})

	//user := &User{
	//	Username: "lisi",
	//	CreditCards: []*CreditCard{
	//		{Number: "0001"},
	//		{Number: "0002"},
	//	},
	//}
	//db.Create(user)
	u := &User{
		Username: "zhangsan",
	}
	db.First(u)
	fmt.Println(u)
	//Assiciation 关联查询,先查user，在查creditcard
	//Association("CreditCards")传进来是个属性名
	db.Model(&u).Association("CreditCards").Find(&u.CreditCards)
	fmt.Println(u.CreditCards)
	userByte, _ := json.Marshal(&u)
	fmt.Println(string(userByte))
	//关联新增
	//在zhangsan下面新增一张信用卡
	//db.Model(&u).Association("CreditCards").Append([]*CreditCard{
	//	{Number: "0003"},
	//})

	//Preload关联查询，先查credit，在查user
	users := []*User{}
	db.Debug().Preload("CreditCards").Find(&users)
	userByte, _ = json.Marshal(users)
	fmt.Println(string(userByte))
}
