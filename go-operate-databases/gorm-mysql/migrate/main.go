package main

// 引入 orm 包时，需要同时引入驱动包
// "gorm.io/driver/mysql" 就是 gorm，mysql 的驱动包
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 表对应的 struct
type User struct {
	ID       int64 `gorm:"primary_key" json:"id"` //sqlx是db
	Username string
	Password string
}

// TableName 自定义表名
func (*User) TableName() string {
	// 返回表名
	return "user_t"
}

// 数据库初始化
func main() {
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	// 建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动创建表, AutoMigrate 没有建立映射关系
	// 用于项目初始化自动创建表
	// 真实项目开发不要这么用,应该由 dba 或者运维定好数据库信息(ddl操作)
	// 默认是 users 结构体名复数
	db.AutoMigrate(User{})
}
