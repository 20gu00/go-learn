package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	t := time.Now()
	fmt.Println(t)

	// 格式化时间 12345
	// time->string
	t2 := t.Format("2006-01-02 15:04:05")
	// fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan")) "# 2021-06-13 13:10:18.143 Sun Jun
	fmt.Println(t2)

	// 时间戳 1970.1.1
	// time->时间戳
	fmt.Println(t.Unix()) // 秒  .UnixNano() 纳秒时间戳 %v

	// 时间戳 -> time
	fmt.Println(time.Unix(t.Unix(), 0))
	fmt.Println(time.Unix(t.Unix(), 0).Format("2006-01-02 15:04:05"))

	// 时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// string->time
	// 格式 string 时区
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", t2, loc)
	fmt.Println(timeObj)
	// time.Parse UTC时间 与time.Now() CST 有8个小时时差
	timeObj2, _ := time.Parse("2006-01-02 15:04:05", t2)
	fmt.Println(timeObj2)
	timeObj3, _ := time.Parse("2006-01-02", t2)
	fmt.Println(timeObj3)

	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	// 打印结果为：2021-05-19 09:20:06
	// 注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute,
		second)

	// time.Duration 是 time 包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
	// time.Duration 表示一段时间间隔，可表示的最长时间段大约290年。

	// Add 时间+时间
	// Sub 两个时间的差值
	now2 := time.Now() // 获取当前时间
	// 1分钟前
	m, _ := time.ParseDuration("-1m") // 1m
	m1 := now2.Add(m)
	fmt.Println(m1)
	// 1分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now2.Add(mm)
	fmt.Println(mm1)

	// 定时器 timer
	// <-NewTimer().C <-time.After()
}

/*2022-12-25 11:19:05.258057849 +0800 CST m=+0.000108603
2022-12-25 11:19:05
1671938345
2022-12-25 11:19:05 +0800 CST
2022-12-25 11:19:05
2022-12-25 11:19:05 +0800 CST
2022-12-25 11:19:05 +0000 UTC
0001-01-01 00:00:00 +0000 UTC
*/
