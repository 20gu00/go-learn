package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter //这个是接口
	R *http.Request       //request是结构体
}

//读取body,处理json数据反序列化
func (c Context) ReadJson(req interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	//反序列化
	err = json.Unmarshal(body, req) //json解码,讲body字节切片处理成我们自己定义的请求
	if err != nil {
		//为什么不写相应呢,因为是出于框架开发者的角度
		//fmt.Fprintf(w, "deserialized failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	//http.ResponceWriter可以调用这些方法,配置responce
	//写入响应头信息
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp) //[]byte
	if err != nil {
		return err
	}
	//写入json数据,经过编码的,[]byte类型
	//这里不用关系写入多少字节的数据
	//记住通信选择默认是json形式,但是会编码用字节的方式传递
	_, err = c.W.Write(respJson)
	return err
}

//这是一种设计,更加方便(当然这样可有可无,看需求,提供便利)就是逐层封装,逐层减少输入,核心是WriteJson方法,可以提供基于它的各种便利的方法

//200
func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

//500
func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

//400
func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

//我们不希望server了解创建context的细节,所以这里提供个函数
func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
