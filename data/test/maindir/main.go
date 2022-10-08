package main

import (
	"fmt"
	"net/http"
	web "tt/v1"
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
	httpServer := web.NewSdkHttpServer("webServer")
	httpServer.Route(http.MethodPost, "/login", Login)
	err := httpServer.Start(":8080")
	if err != nil {
		panic(err) //server启动失败的话直接报错panic就好
	}
}

func Login(ctx *web.Context) {
	req := &loginReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &loginResponse{
		Data: "不同全部字段都写",
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		ctx.BadRequestJson("响应失败")
		//fmt.Printf("响应失败:%v", err)
	}
}

type loginReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type loginResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
