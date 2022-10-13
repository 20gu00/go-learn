package web

import (
	"net/http"
	"strings"
)

type HandlerBaseTree struct {
	root *node
}

func NewHandlerBaseTree() Handler {
	return &HandlerBaseTree{
		root: &node{},
	}
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 10),
	}
}
func (h *HandlerBaseTree) ServeHTTP(c *Context) {
	handler, found := h.searchRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("not found!"))
		return
	}

	handler(c)
}

func (h *HandlerBaseTree) searchRouter(pattern string) (handlerFunc, bool) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	currentRoot := h.root
	for _, path := range paths {
		child, ok := h.searchChild(currentRoot, path)
		if !ok {
			return nil, false
		}
		currentRoot = child
	}
	if currentRoot.handler == nil {
		return nil, false
	}
	return currentRoot.handler, true
}

func (h *HandlerBaseTree) HttpRoute(
	method string,
	pattern string,
	handlerFunc handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	currentRoot := h.root
	for index, path := range paths {
		matchChild, ok := h.searchChild(currentRoot, path)
		if ok {
			currentRoot = matchChild
		} else {
			// 为当前节点根据
			h.buildSubTree(currentRoot, paths[index:], handlerFunc)
			return
		}
	}
	currentRoot.handler = handlerFunc
}

func (h *HandlerBaseTree) searchChild(root *node, path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func (h *HandlerBaseTree) buildSubTree(root *node, paths []string, handlerFn handlerFunc) {
	currentRoot := root
	for _, path := range paths {
		newN := newNode(path)
		currentRoot.children = append(currentRoot.children, newN)
		currentRoot = newN
	}
	currentRoot.handler = handlerFn
}
