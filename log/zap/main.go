package main

import (
	"go.uber.org/zap"
	"net/http"
)

var Logger *zap.Logger
var SugerLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer Logger.Sync() //同步日志,刷新到磁盘上

	url := "www.badiu.com"  //一般得加上http://才正确
	_, err := http.Get(url) //也可以打印resp
	if err != nil {
		//msg错误描述string zapFiled键值对格式显示
		//msg:xxx url:xxx error:xxx
		Logger.Error("错误描述", zap.String("key", url), zap.Error(err))
	}

	//默认json格式,输出到终端

	//trace debug warn info error panic fatal
	//Logger.Info
	//Logger.Error("错误描述",zap.String("key",v),zap.Error(err))
}

func InitLogger() {
	Logger, _ = zap.NewProduction()

	//如果要的是SugarLogger
	SugerLogger = Logger.Sugar()

	//生产级别日志不容易看出区别,NewDevelopment  每个字段一个空格显示
	//SugerLogger.Info
	//SugerLogger.Error("错误描述",zap.String("key",v),zap.Error(err))
}
