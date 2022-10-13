package web

import "net/http"

type HttpServer interface {
	RouteDo
	ServerStart(address string) error
}

type RouteDo interface {
	HttpRoute(method string, pattern string, handlerFunc handlerFunc)
}
type factServer struct {
	Name    string
	root    Filter
	handler Handler
}

func (f *factServer) ServerStart(addr string) error {
	//根
	http.HandleFunc("/", func(writer http.ResponseWriter,
		request *http.Request) {
		c := NewContext(writer, request)
		f.root(c)
	})
	return http.ListenAndServe(addr, nil)
}

func (f *factServer) HttpRoute(method string, pattern string,
	handlerFunc handlerFunc) {
	f.handler.HttpRoute(method, pattern, handlerFunc)
}

func NewFactServer(name string, builders ...FilterBuilder) HttpServer {
	handler := NewHandlerBaseTree()
	//handler := NewHandlerBasedOnMap()
	// 因为我们是一个链，所以我们把最后的业务逻辑处理，也作为一环
	var root Filter = handler.ServeHTTP
	// 从后往前把filter串起来
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	res := &factServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
	return res
}
