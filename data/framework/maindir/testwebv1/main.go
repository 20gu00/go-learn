package main

import (
	"fmt"
	web "github.com/20gu00/web1/v1"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "用户页面")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "添加用户")
}

func orderPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "订单页面")
}

func main() {
	httpServer := web.NewFactServer("webServer")
	httpServer.HttpRoute(http.MethodPost, "/login", Login)
	err := httpServer.ServerStart(":8080")
	if err != nil {
		panic(err) //server启动失败的话直接报错panic就好
	}
}

func Login(ctx *web.Context) {
	req := &loginReq{}

	err := ctx.GetJson(req)
	if err != nil {
		ctx.ErrorRequestJson(err)
		return
	}

	resp := &loginResponse{
		Data: "不同全部字段都写" + req.Email,
	}
	err = ctx.SentJson(http.StatusOK, resp)
	if err != nil {
		ctx.ErrorRequestJson("响应失败")
		//fmt.Printf("响应失败:%v", err)
	}
}

type loginReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type loginResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
