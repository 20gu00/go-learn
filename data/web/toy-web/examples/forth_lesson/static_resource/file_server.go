package main

import "net/http"

func main() {
	serve := http.FileServer(http.Dir(".")) //toy-web下边的内容,包括隐藏文件可以实现个简单的文件服务器
	//http.Handle("/", serve)
	http.ListenAndServe(":8080", serve)
}

//也就是整个项目目录,可以以go.mod为参考
