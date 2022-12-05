package main

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	//默认值
	//viper.SetDefault("k1", "v1")

	//读取配置文件
	//配置文件名
	viper.SetConfigName("base") //可以不带后缀($HOME目录下的比如.bashrc .bash_profile直接写)
	//配置文件类型
	viper.SetConfigType("yaml")
	//查找配置文件的路径(可以重复调用,设置多个搜索路径)
	viper.AddConfigPath(".") // ./
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("未找到配置文件")
		} else {
			fmt.Println("读取配置文件失败")
			panic(errors.New("读取配置文件失败"))
		}
		//panic(err)
	}

	////写入配置文件
	////将当前的配置写入预定义的路径,覆盖,如果没有预定义路径(path configName)则会报错
	//viper.WriteConfig()  //会创建
	////不同之处就是文件已经存在,不会覆盖
	//viper.SafeWriteConfig()
	//
	////将当前的viper配置写入配置文件,覆盖
	//viper.WriteConfigAs("./file")
	////不同之处不会覆盖
	//viper.SafeWriteConfigAs("./file")
	//
	//热加载,运行时监控配置文件
	viper.WatchConfig() //前提是已经添加设置好了配置文件(ReadInConfig)
	//配置文件变更后的回调函数(可选)
	viper.OnConfigChange(func(e fsnotify.Event) {
		//回调逻辑
		fmt.Println("配置文件已经修改")
	},
	)

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		//最好只有一个string参数作format
		c.String(http.StatusOK, viper.GetString("ver.v"))
	})
	r.Run()
}
