package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RDB *redis.Client

func main() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	defer RDB.Close()

	if _, err := RDB.Ping().Result(); err != nil {
		fmt.Println("reids建立连接失败", err)
		return
	}

	fmt.Println("建立redis的连接成功")

	key := "watch_counter" //设置key
	err := RDB.Watch(func(tx *redis.Tx) error {
		//watch这个key
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil { //不是空,空也是一种错误
			return err //Get操作失败
		}
		//key加1
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			//业务逻辑
			//并不能完全保证数据的冲突性,比如服务端修改值,后台管理员修改了价格
			//time.Sleep(5*time.Second)  这期间服务器修改值
			pipe.Set(key, n+1, 0) //不会过期
			return nil            //这里不处理错误逻辑
		})
		return err
	}, key)
	if err != nil {
		fmt.Println("事务执行失败", err)
		return
	}

	fmt.Println("事务执行成功")
}
