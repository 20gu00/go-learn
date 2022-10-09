package web

import (
	"net/http"
	"sync"
)

var _ Handler = &HandlerBaseMap{} //判断HandlerBaseMap是否是Handler类型,这里其实就是HandlerBaseMap是否实现了Handler接口

type HandlerBaseMap struct {
	handlers sync.Map //线程安全的map
}

func NewHandlerBaseMap() *HandlerBaseMap {
	//return &HandlerBaseOnMap{
	//	handlers: make(map[string]func(c *Context), 100), //一般一个用户注册的路由不会太多,给100个这里
	//}

	//直接用线程安全的map,而不是原生的map
	return &HandlerBaseMap{}
}

func (h *HandlerBaseMap) ServeHTTP(c *Context) {
	req := c.R
	key := h.getKey(req.Method, req.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("未匹配到相应的路由"))
		return
	}

	handler.(handlerFunc)(c)
}

func (h *HandlerBaseMap) HttpRoute(method string, pattern string,
	handlerFunc handlerFunc) {
	key := h.getKey(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func (h *HandlerBaseMap) getKey(method string,
	path string) string {
	return method + "@" + path
}
