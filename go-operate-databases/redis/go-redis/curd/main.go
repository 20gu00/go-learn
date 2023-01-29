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

	// 过期时间
	if err := RDB.Set("score", 100, 0).Err(); err != nil {
		fmt.Println("set失败", err.Error())
		return
	}
	//.Val()方法不会返回错误,如果有错误就是类型的零值
	value, err := RDB.Get("score").Result()
	if err != nil {
		fmt.Println("get失败", err.Error())
		return
	}

	fmt.Println(value)

	value2, err := RDB.Get("score").Result()
	if err == redis.Nil {
		fmt.Println("score key not exist")
	} else if err != nil {
		fmt.Println("获取score的值失败")
	} else {
		fmt.Println(value2)
	}

	//hash value是map
	//RDB.HGetAll() HGet() HMGet()
	//hset user name "aa"  map类型 user(key) value:{"name":"aa"}
	//hset user age 100
	//hgetall user
	//redis-cli

	//获取这个key下所有的value
	value3, err := RDB.HGetAll("user").Result()
	if err == redis.Nil {
		fmt.Println("user key not exist")
	} else if err != nil {
		fmt.Println("user")
	} else {
		fmt.Println(value3)
	}

	//获取该key的value的指定的key的值
	RDB.HMGet("user", "name", "age").Result()
	RDB.HGet("user", "name").Val()

	//zset 排行榜 交集 有序集合
	zsetKey := "language_rank"
	language := []redis.Z{
		//Score Member
		redis.Z{Score: 90.0, Member: "go"},
		redis.Z{Score: 98.0, Member: "java"},
		redis.Z{Score: 95.0, Member: "python"},
		redis.Z{Score: 97.0, Member: "js"},
		redis.Z{Score: 99.0, Member: "c"},
	}

	//zadd
	num, err := RDB.ZAdd(zsetKey, language...).Result() //要传递的参数是变长(可变)参数
	if err != nil {
		fmt.Println("zadd failed", err.Error())
		return
	}

	fmt.Println(num)

	//increase
	newScore, err := RDB.ZIncrBy(zsetKey, 10.0, "go").Result()
	if err != nil {
		fmt.Println("增加失败", err.Error())
		return
	}

	fmt.Println(newScore)
	//取分数最高的三个
	ret, err := RDB.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrange失败", err.Error())
		return
	}

	for _, v := range ret {
		fmt.Println(v.Member, v.Score)
	}

	//取出95-100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = RDB.ZRevRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Println("获取95-100分的成员失败")
		return
	}

	for _, v := range ret {
		fmt.Println(v.Member, v.Score)
	}

	//列出无顺序
	//zrevrange language_rank 0 2
	//zrevrange language_rank 0 2 withscores

	//列出有顺序
	//zrange language_rank 0 2  从大到小
	//zrange language_rank 0 2 withscores

	//zincrby language_rank 10 "go"
}
