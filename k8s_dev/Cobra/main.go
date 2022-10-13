/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS> 版权所有。

根据Apache许可证，2.0版（"许可证"）许可。
除非遵守许可证的规定，否则你不得使用此文件。
你可以在以下网址获得一份许可证的副本

    http://www.apache.org/licenses/LICENSE-2.0

除非适用法律要求或书面同意，根据本许可证分发的软件
许可证下分发的软件是以 "原样 "为基础的。
没有任何形式的保证或条件，无论是明示还是暗示。
有关许可证下的权限和限制的具体语言，请参见许可证。
许可证下的许可和限制的具体语言。
*/

package main

import "kubefetch/cmd"

//init->main->initConfig->rootCmd的Run
func main() {
	cmd.Execute()
}
