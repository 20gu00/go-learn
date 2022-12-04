package main

import (
	"log"
	"os"
)

func main() {
	logFileLocation, _ := os.OpenFile("/root/pull/go-learn/log/default-go-logger/go.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744) //追加
	log.SetOutput(logFileLocation)                                                                                                 //os.Stdout
	log.Printf("a")                                                                                                                //只有Print,不支持级别
	log.Fatal("a")                                                                                                                 //程序直接退出,defer也不执行,不正常退出
}

//log.Printf
