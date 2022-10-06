package main

import "net/http"

type Routable interface { //可路由的
	Route(method string,
		pattern string,
		//handlerFunc func(ctx *Context))
		handlerFunc handlerFunc)
}

type Handler interface {
	//也就成了一方面处理请求,一方面怎么路由

	//我把context封装好了,用这个它有将context拆开,所以不用这个Handler
	//http.Handler

	ServeHTTP(c *Context)
	Routable
	//Route(method string, pattern string, handlerFunc func(ctx *Context))
}

//基于map的handler
//也就是用map来记录路由注册,但原生的map有线程安全问题,方法一是你强制用户一开始就注册号路由(单一线程),运行中再注册路由就报错,就是注册玩路由在启动server方法二就是自己实现线程安全,允许用户随意出注册路由
//所有的同步都是有代价的,性能问题,sync性能远高于sync.map
type HandlerBaseOnMap struct {
	//所有的handlers
	//key应该是method+url
	handlers map[string]func(c *Context)
}

func (h *HandlerBaseOnMap) Route(
	method string,
	pattern string,
	handleFunc handlerFunc) {
	//注意这样其实并没有解决重复注册
	key := h.Key(method, pattern)
	h.handlers[key] = handleFunc
}

//拥有多个filter,有handler调用
//自定义的handler
func (h *HandlerBaseOnMap) ServeHTTP(c *Context) { //(writer http.ResponseWriter, request *http.Request) {
	//panic("hi hi")
	key := h.Key(c.R.Method, c.R.URL.Path)
	//判断该路由是否已经注册过
	//如果key不存在comma,ok值为false,第一个值是零值
	if handler, ok := h.handlers[key]; ok {
		//已经找到的就是注册过的,调用一下
		handler(c) //就不需要再New个context了,已经有了(NewContext(writer, request))
	} else {
		//没找到,就是该路由没有注册,返回个404
		//响应头
		c.W.WriteHeader(http.StatusNotFound)
		//响应的信息
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBaseOnMap) Key(method string, pattern string) string {
	//参数request *http.Request,path是肯定会拿到的
	//return req.Method + "#" + req.URL.Path
	return method + "#" + pattern
}

var _ Handler = &HandlerBaseOnMap{}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(c *Context), 100), //一般一个用户注册的路由不会太多,给100个这里
	}
}
