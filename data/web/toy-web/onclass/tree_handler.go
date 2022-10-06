package main

import (
	"net/http"
	"strings"
)

type HandlerBaseOnTree struct {
	root *node //根节点
}
type node struct {
	path     string
	children []*node //根节点中有个children,里边是它的所有子节点

	//如果这是叶子节点,那么匹配上后就可以调用该方法
	handler handlerFunc
}

//HTTP Server: 新增一条路由
/*
新增/user/friend步骤:
1.从根节点出发(/),作为当前节点-->/
2.查找命中的子节点
3.讲子节点作为当前节点-->/user  然后重复2
4.如果当前节点没有匹配下一段(就是说当前节点找不到friend),创建子节点
5.如果找到了也就是说路径还没结束,那就继续2 3
6.创建成功
*/

func (h *HandlerBaseOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

//这里是查找路由,很简单
func (h *HandlerBaseOnTree) findRouter(path string) (handlerFunc, bool) {
	// 去除头尾可能有的/，然后按照/切割成段
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前 path 的节点
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	// 到这里，应该是找完了
	if cur.handler == nil {
		// 到达这里是因为这种场景
		// 比如说你注册了 /user/profile
		// 然后你访问 /user
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBaseOnTree) Route(method string,
	pattern string,
	handlerFunc handlerFunc) {
	// /user/friend/ --> ["","user","friend",""]  /user/friend --> ["","user","friend"]
	//先将前后的/去掉  默认是去除空格
	pattern = strings.Trim(pattern, "/")
	//切割 /user/friend --> paths:["user","friend"]
	paths := strings.Split(pattern, "/")

	//从root出发
	cur := h.root

	//遍历path
	for index, path := range paths {
		//自顶向下写法,即找到子节点
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			//如果找到了,就替换"根节点",即用当前节点
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handlerFunc) //注意是创建子树,而不是节点,因为用户要创建的路径可能很长,含有多个子节点
			return
		}
	}
}

//这就是完全匹配.严格匹配的模式,不够灵活,必须完全匹配整个路径,不支持*
//从根节点找path,当然根节点即当前节点,可以替换,就是使当前节点
//其实这里接受者传*node会更好,那么就不用再传递个形参root *node
//思路逻辑就是我自己知道怎么找子节点,而不是用我的人知道怎么找子节点
func (h *HandlerBaseOnTree) findMatchChild(root *node, path string) (*node, bool) {
	//就是遍历全部的注册好的子节点(叶子)
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

//还要加上叶子节点的业务处理逻辑
func (h *HandlerBaseOnTree) createSubTree(root *node, paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		//就是向当前节点中不断append进子节点
		cur.children = append(cur.children, nn)
		cur = nn
	}
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2), //要初始化好,后续要append,make会初始化
	}
}
