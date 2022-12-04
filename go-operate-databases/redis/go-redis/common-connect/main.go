package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RDB *redis.Client //"github.com/go-redis/redis/v8"

func main() {
	//自己会维护一套连接池
	//注意没有冒号,尤其是作为调包使用的时候,初始化的是全局变量
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //"locahost:6379"
		Password: "",
		DB:       0,   //使用默认的db(使用哪个库)
		PoolSize: 100, //连接池大小
	})

	//记得释放
	defer RDB.Close() //如果是做成包调用,记得在外部defer释放

	//建立连接(原生的一般都是ping建立连接和测试连通性
	if _, err := RDB.Ping().Result(); err != nil {
		fmt.Println("reids建立连接失败", err)
		return
	}

	fmt.Println("建立redis的连接成功")
}
