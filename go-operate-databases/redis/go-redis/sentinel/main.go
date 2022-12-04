package main

import "github.com/go-redis/redis"

var RDB *redis.Client //哨兵模式依旧是client

func main() {
	RDB = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"127.0.0.1:6379", "192.168.23.10:6379", "192.168.23.239:6379"},
	})
	if _, err := RDB.Ping().Result(); err != nil {
		return
	}

}
