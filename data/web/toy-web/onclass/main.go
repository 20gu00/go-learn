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
	//可以看到这种API不限制方法Get Put Delete Post
	//Server.Route("/user/create", createUser)
	Server.Route(http.MethodGet, "/user/signup", SignUp)
	//Server.Route("/order", order)

	//Server.Route(":8080", nil)
	err := Server.Start(":8080")
	if err != nil {
		panic(err)
	}
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

//错误error是个interface,小写但是能包外访问,在buildin文件中,建议优先考虑error而不是panic
//遇事不决用指针
//快速失败的场景比如服务器启动失败,这种的就直接用panic
//一般是自己觉得,至少自己觉得这个地方的错误"不可恢复",所以写panic,可以这么理解自己又写panic又revocer就很奇怪,一般都是recover别人写的代码
