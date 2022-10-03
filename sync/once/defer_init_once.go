package main

import (
	"fmt"
	"net"
	"sync"
)

// 使用连接
func main() {
	var addr = "baidu.com:80" //注意新版本要指定端口号
	var conn net.Conn
	var once sync.Once
	once.Do(func() {
		conn, _ = net.Dial("tcp", addr)
	})
	if conn == nil {
		panic("conn is nil")
	}

	fmt.Println("ok")
}
