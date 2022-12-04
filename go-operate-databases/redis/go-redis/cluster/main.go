package main

import "github.com/go-redis/redis"

var RDB *redis.ClusterClient //ClusterClient

func main() {
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":9000", ":9001", "9002"},
	})
	if _, err := RDB.Ping().Result(); err != nil {
		return
	}

}
