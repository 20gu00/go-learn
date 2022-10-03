package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	copyFile3("/root/go-learn/io/file_copy/test1", "/root/go-learn/io/file_copy/test2")
}

//第三种方法是使用ioutil包中的 ioutil.WriteFile()和 ioutil.ReadFile()，但由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出。
func copyFile3(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("操作失败：", destFile)
		fmt.Println(err)
		return 0, err
	}

	return len(input), nil
}
