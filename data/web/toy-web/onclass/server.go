package main

import (
	"encoding/json"
	"fmt"
	webv2 "geektime/toy-web/pkg/v2"
	"net/http"
)

type Server interface {
	//这是手动创建context的期望
	//Route(pattern string, handleFunc http.HandlerFunc) //handler func(ResponseWriter, *Request)

	//这个是由框架生成context
	//Route(pattern string, handlerFunc func(ctx *Context))

	//Restfule 风格(其实就是原生库加多个方法)
	//method+pattern指定一个操作的资源路径

	//Route(method string, pattern string, handlerFunc func(ctx *Context))
	Routable
	Start(addr string) error
}

//基于http库实现(也可以是第三方库)
type sdkHttpServer struct {
	Name    string //大写也没用
	handler Handler
	root    Filter
}

//Route路由注册
//请求路径和处理逻辑方法方法绑定
func (s *sdkHttpServer) Route(
	method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	s.handler.Route(method, pattern, handleFunc)
	//注意这样其实并没有解决重复注册
	//key := s.handler.Key(method, pattern)
	//s.handler.handlers[key] = handleFunc

	//restful的案例,因为原生的http.HandleFunc()不支持method参数
	//因为只注册一遍,所以挪到Start那里
	//http.Handle("/", &HandlerBaseOnMap{}) //第二个参数是interface,我们这里用了自定义的结构体,所以该结构体要实现该接口

	//这里是框架生成context的案例
	////闭包,调用,匿名
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) { //注意这里只是函数的形参
	//	//函数体
	//
	//	//在这里创建context
	//	//ctx:=&Context{
	//	//	R:request,
	//	//	W:writer,
	//	//}
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
}

func (s *sdkHttpServer) Start(addr string) error {
	//http.Handle("/", &HandlerBaseOnMap{})
	//handler:=&HandlerBaseOnMap{}

	//用filter了就不是直接handler
	//http.Handle("/", s.handler)  //server最终是委托给handler,那就看handler文件

	http.Handle("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c) //已经做好了filter链,root里链会自动调用下一个链,当然链是从后往前拼的
		//比如说这里root是MetricsFilterBuilder,那么next就是下一个处理请求
	})
	return http.ListenAndServe(addr, nil)
}

//返回接口对象
func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBaseOnMap()
	//定义个root的filter
	var root Filter = func(c *Context) {
		handler.ServeHTTP(c.W, c.R)
	}
	//从后往前调用next
	for i := len(builders); i >= 0; i-- {
		//从后往前拼成链
		b := builders[i]
		root = b(root) //第一个filter就是root这个filter,真正处理请求的filter,如果不传filter就是直接处理请求
		//当传入filter,就会当做是next调用前一个filter,从后往前串联起来
	}
	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	} //注意返回而是指针,返回实现该接口的对象
	//return factory() //工厂模式,一种设计模式
}

////工厂模式更好,方便多个接口
//type Factory func() Server
//
//var factory Factory
//
//func RegisterFactory(f Factory) {
//	factory = f
//}

type signUpReq struct {
	Email             string `json:"email"` //tag,可以运行时通过反射拿到
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//使用框架提供的context,不用自己创建
//所以框架开发者希望你在绑定路由是传递处理函数是func(ctx *Context)这种形式
func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		//这里一般要写入日志
		fmt.Printf("写入相应失败\n")
	}
}

// 在没有 context 抽象的情况下，是长这样的
//func SignUp(w http.ResponseWriter, r *http.Request) {
//	req := &signUpReq{}
//
//	//有个缺陷就是用户手动创建context,这样框架开发者就不好控制,我们希望由开发者来控制contex,也就是上下文以来框架自身去创建
//	ctx := &Context{
//		W: w,
//		R: r,
//	}
//	err := ctx.ReadJson(req)
//	if err != nil {
//		fmt.Fprintf(w, "err: %v", err) //%v会自动解析成go的表达
//		return
//	}
//
//	resp := &commonResponse{
//		Data: 123,
//	}
//	err=ctx.WriteJson(http.StatusOK,resp)
//	if err !=nil {
//		//这里一般要写入日志
//		fmt.Printf("写入相应失败\n")
//	}
//
//	//body, err := io.ReadAll(r.Body)
//	//if err != nil {
//	//	fmt.Fprintf(w, "read body failed: %v", err)
//	//	// 要返回掉，不然就会继续执行后面的代码
//	//	return
//	//}
//	////反序列化
//	//err = json.Unmarshal(body, req)  //json解码,讲body字节切片处理成我们自己定义的请求
//	//if err != nil {
//	//	fmt.Fprintf(w, "deserialized failed: %v", err)
//	//	// 要返回掉，不然就会继续执行后面的代码
//	//	return
//	//}
//
//	// 返回一个虚拟的 user id 表示注册成功了
//	//fmt.Fprintf(w, "%d", 123)
//
//	//应该返回个json
//	//fmt.Fprintf(w, `{"userId:": %d}`, 123)
//
//	//例子
//	//resp := &commonResponse{
//	//	Data: 123,
//	//}
//
//	//respJson, err := json.Marshal(resp) //[]byte
//	//if err != nil {
//	//
//	//}
//	////返回个响应头
//	//w.WriteHeader(http.StatusOK)
//	//fmt.Fprintf(w, string(respJson))
//
//	//可见很多重复的工作,像用Server来进一层封装同样的思路,我们用context上下文来替代请求的信息的传递,即上面这些
//}

func SignUpWithoutWrite(w http.ResponseWriter, r *http.Request) {
	c := webv2.NewContext(w, r)
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		resp := &commonResponse{
			BizCode: 4, // 假如说我们这个代表输入参数错误
			Msg:     fmt.Sprintf("invalid request: %v", err),
		}
		respBytes, _ := json.Marshal(resp)
		fmt.Fprint(w, string(respBytes))
		return
	}
	// 这里又得来一遍 resp 转json
	fmt.Fprintf(w, "invalid request: %v", err)
}
