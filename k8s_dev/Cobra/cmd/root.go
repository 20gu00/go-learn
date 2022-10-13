/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS> 版权所有。

根据Apache许可证，2.0版（"许可证"）许可。
除非遵守许可证的规定，否则你不得使用此文件。
你可以在以下网址获得一份许可证的副本

    http://www.apache.org/licenses/LICENSE-2.0

除非适用的法律要求或以书面形式同意，根据本许可证分发的软件
许可证下分发的软件是以 "原样 "为基础的。
没有任何形式的保证或条件，无论是明示还是暗示。
有关许可证下的权限和限制的具体语言，请参见许可证。
许可证下的许可和限制的具体语言。
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper" //优秀的处理配置文件的库,json homl yaml等等
)

var cfgFile string

// rootCmd代表在没有任何子命令的情况下调用的基本命令。
var rootCmd = &cobra.Command{
	Use:   "kubefetch",
	Short: "A brief description of your application", //对你的申请的简要描述
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`, /*一个较长的描述，横跨多行，可能包含
	使用你的应用程序的例子和用法。比如说。

	Cobra是一个Go的CLI库，为应用程序赋能。
	这个应用程序是一个生成所需文件的工具
	来快速创建一个Cobra应用程序。*/
	// 如果你的裸应用程序有一个与之相关的动作，请取消下面这一行的注释。
	// 有一个与之相关的动作。
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	},
}

// Execute将所有子命令添加到根命令中，并适当地设置标志。
// 这是由main.main()调用的。它只需要发生在rootCmd上一次。
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() { //init函数导包时就调用,常量变量init函数
	cobra.OnInitialize(initConfig)

	// 在这里你将定义你的标志和配置设置。
	// Cobra支持持久性标志，如果在这里定义了这些标志。
	// 将成为你的应用程序的全局。

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubefetch.yaml)")

	// Cobra也支持本地标志，只有在直接调用这个动作时才会运行
	// 当这个动作被直接调用时。
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig 读取配置文件和 ENV 变量（如果设置）。
//先执行initConfig再到rootCmd的Run
func initConfig() {
	if cfgFile != "" {
		// 使用flag上的配置文件。
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找主目录。
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 在主目录中搜索配置，名称为".kubefetch"（无扩展名）。
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".kubefetch")
	}

	viper.AutomaticEnv() // 读取符合以下条件的环境变量

	// 如果发现了一个配置文件，就把它读进去。
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
