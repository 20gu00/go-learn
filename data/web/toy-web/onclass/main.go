package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	//自己封装路由注册
	Server := NewHttpServer("test-server")
	//Server.Route("/", home)
	//Server.Route("/user", user)
	//Server.Route("/user/create", createUser)
	Server.Route("/user/signup", SignUp)
	//Server.Route("/order", order)
	Server.Route(":8080", nil)
	Server.Start(":8080")

	//http.HandleFunc("/", home)
	//http.HandleFunc("/user", user)
	//http.HandleFunc("/user/create", createUser)
	//http.HandleFunc("/order", order)
	//http.ListenAndServe(":8080", nil)
}

//type Server interface {
//	Route(pattern string, handlerFunc http.HandlerFunc)
//	Start(address string) error
//}

//type sdkHttpServer struct {
//	Name string
//}
