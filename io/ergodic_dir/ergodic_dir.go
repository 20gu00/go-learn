package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/**
	  遍历文件夹：
	*/
	//学习io之后，尤其是文件操作，我们就可以遍历给定的目录文件夹了。可以使用ioutil包下的readDir()方法，这个方法可以获取指定目录下的内容，返回文件和子目录。
	//
	//因为文件夹下还有子文件夹，而ioutil包的ReadDir()只能获取一层目录，所以我们需要自己去设计算法来实现，最容易实现的思路就是使用递归。
	dirname := "." //即当前文件夹,注意工程目录,稳妥的话可以字节Chdir噶边工作目录,或者使用绝对路径
	listFiles(dirname, 0)

}

func listFiles(dirname string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename, level+1)
		}
	}
}
