package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//`gorm:"many2many:user_languages"` 该tag用于声明多对多关系，以及中间表的名字
type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages"`
}

// 自定义第三张表
// gorm.Model有主键
type UserLanguage struct {
	gorm.Model
	UserID     int `gorm:"primarykey"`
	LanguageID int `gorm:"primarykey"`
	//////////////
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//自动创建表,AutoMigrate没有建立映射关系
	//用于项目初始化自动创建表
	//中间表
	db.AutoMigrate(User{}, Language{}, UserLanguage{})

	//添加数据
	//user := &User{
	//	Languages: []*Language{
	//		{Name: "go"},
	//		{Name: "vue"},
	//	},
	//}
	//db.Create(&user)

	//关联查询

	//预加载Preload获取数据
	users := []*User{}

	db.Preload("Languages").Find(&users)
	// 对这个[]json编码
	userByte, _ := json.Marshal(users)
	fmt.Println(string(userByte))
	//Preload获取单条数据
	user := &User{
		Model: gorm.Model{
			ID: 1,
		},
	}
	//User.Languages
	db.Preload("Languages").Find(&user)
	fmt.Println(user)
	//Association获取单条数据
	db.Model(&user).Association("Languages").Find(&user.Languages)
	fmt.Println(user)
}
