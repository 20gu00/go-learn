package main

import (
	"fmt"
	"log"
	"net/http"
)

//对应该url的处理逻辑函数,url是完整的请求,即ip或dns和port和path
//又叫路由
func handler1(w http.ResponseWriter, r *http.Request) {
	//将内容写入到w中,Fprintf的第一个参数是能实现write方法的接口,http.ResponseWriter接口包含了该方法
	//write方法是用来写入数据的
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //要的是path,1开始即不要第一个/
}

func handler2(w http.ResponseWriter, r *http.Request) {
	//将内容写入到w中,Fprintf的第一个参数是能实现write方法的接口,http.ResponseWriter接口包含了该方法
	//write方法是用来写入数据的
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //要的是path,1开始即不要第一个/
}
func main() {
	//讲请求的路径和处理函数绑定
	http.HandleFunc("/", handler1)
	http.HandleFunc("/a", handler2)
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //要的是path,1开始即不要第一个/

	})

	//监听,提供服务,指定端口
	log.Fatal(http.ListenAndServe(":8080", nil))
}
