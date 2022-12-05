package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//覆盖
	viper.Set("k1", "v1")

	//默认值
	viper.SetDefault("k1", "v1")

	//读取配置文件
	//配置文件名
	viper.SetConfigName("base") //可以不带后缀($HOME目录下的比如.bashrc .bash_profile直接写)
	//配置文件类型
	viper.SetConfigType("yaml")
	//查找配置文件的路径(可以重复调用,设置多个搜索路径)
	viper.AddConfigPath(".") //not./
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("未找到配置文件")
		} else {
			panic(errors.New("读取配置文件失败"))
		}
		//panic(err)
	}

	//写入配置文件
	//要注意文件的覆盖
	//将当前的配置写入预定义的路径,覆盖,如果没有预定义路径(path configName)则会报错
	viper.WriteConfig()
	//不同之处就是文件已经存在,不会覆盖
	viper.SafeWriteConfig()

	//一般没有都会自动创建
	//将当前的viper配置写入配置文件,覆盖
	viper.WriteConfigAs("./file1")
	//不同之处不会覆盖
	viper.SafeWriteConfigAs("./file1")

	//热加载,运行时监控配置文件
	viper.WatchConfig() //前提是已经添加设置好了配置文件(ReadInConfig)
	//配置文件变更后的回调函数(可选)
	viper.OnConfigChange(func(e fsnotify.Event) {
		//回调逻辑
		fmt.Println("配置文件已经修改")
	},
	)

	//自定义配置源 io.Reader(buffer)

	viper.SetConfigType("yaml")
	//io处理的都是[]byte
	//string->[]byte
	var yaml = []byte(`
		k1: v1
		k2:
		- a
		- b
		k3: v
		`)

	viper.ReadConfig(bytes.NewBuffer(yaml))
	fmt.Println(viper.Get("k1"))

	//注册和使用别名
	//两个key绑定一块,get set操作任意一个等于操作了两个
	viper.RegisterAlias("k1", "k2")

	//使用环境变量
	//区分大小写
	//viper自身机制环境变量唯一
	//SetEnvPrefix设置前缀,BindEnv和AutomaticEnv都是用这个前缀(比如之先处理http开头的环境变量)

	//BindEnv()一个或者两个参数,key env_name
	//如果不提供env_name,viper设置环境变量:前缀_key大写
	//如果提供比如id,viper会去查找ID,不会自动添加前缀

	//AutomaticEnv viper.Get()时会检查环境变量
	//如果设置了EnvPrefix:
	//SetEnvKeyReplacer语序使用strings.Replacer会一定程度重写ENV(键),比如Get()时想用-或其他符号,但是env是_

	//默认情况下空环境变量被认为是未设置的,返回下一个配置源,AloowEmptryEnv设置空的也是设置的

	viper.SetEnvPrefix("aaa") //自动AAA
	viper.BindEnv("bbb")      //找BBB环境变量

	//一般不再代码中设置
	os.Setenv("A", "a")
	id := viper.Get("id") //环境变量ID
	fmt.Println(id)

	//使用flags
	//viper使用的是cobra的pflag
	//serverCmd.Flags().Int("port", 9000, "describe")
	//viper.BindPFlag("port", serverCmd.Flags().Lookup("Port"))

	//或者标准的flag库

	//一组绑定
	pflag.Int("name", 1000, "describe")
	//可以换成flag.Int("name",1000,"describe")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.GetInt("name")

	//远程key/value存储
	//import _ "github.com/sqf13/viper/remote"
	//详情看github https://github.com/spf13/viper

	//设置好了viper,读取了,那就可以获取值
	//viper的每一个Get方法获取不到时会返回零值,可以通过IsSet()判断(先判断)
	//不区分大小写(viper基本上就只有设置和获取ENV时需要注意大小写)
	//嵌套 .
	//如果key中含有.,县一个一个找,找不到满足的在找完整key

	//获取子树
	//a:
	//  b:
	//    aa
	//    AA

	//viper.Sub("a.b")

	//aa
	//AA

	//序列化和反序列化
	//viper获取的数据和结构体(一般都是结构体)
	//viper.Unmarshal(&Struct)
	c := new(Config)
	viper.Unmarshal(c) //err
	fmt.Printf("c:%#v", c)

	//如果解析的key包含:  这是viper的key默认的分割符号
	//设置key的分割符
	viper.NewWithOptions(viper.KeyDelimiter("::"))

	//yaml->string 序列化
	//yaml "gopkg.in/yaml.v2"
	//c:=viper.AllSetting()
	//bs,err:=yaml.Marshal(c)

	//多个viper使用,就多个viper.New()
}

type Config struct {
	Ka          string `mapstructure:"ka"` //不论配置文件什么格式都要配置这个,不然操作不了这个字段
	Kb          string `mapstructure:"kb"`
	MysqlConfig `mapstructure:"mysql"` //嵌套匿名结构体,模拟了OOP的继承
}
type MysqlConfig struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	Version string `mapstructure:"version"`
	Db      string `mapstructure:"db"`
}
