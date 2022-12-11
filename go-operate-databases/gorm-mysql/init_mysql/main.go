package main

//引入orm包时，需要同时引入驱动包
//"gorm.io/driver/mysql"就是gorm，mysql的驱动包
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库初始化
func main() {
	//定义连接地址
	//parseTime=True&loc=Local
	//parseTime是查询结果是否自动解析为时间(数据库datetime或者timestamp <=>go的time.Time)
	//id name type  create_time(time.Time)  数据库字段格式datetime
	//loc是Mysql的时区设置
	dsn := "root:100.Acjq@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local" //"Asia/Shanghai"
	//建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
