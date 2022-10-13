package web

type handlerFunc func(c *Context)

type Handler interface {
	RouteDo
	ServeHTTP(c *Context)
}
