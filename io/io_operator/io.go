package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   读取数据：
	       Reader接口：
	           Read(p []byte)(n int, error)
	*/
	//读取本地aa.txt文件中的数据
	//step1：打开文件
	fileName := "/root/go-learn/io/io_operator/test.file"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//step3：关闭文件
	defer file.Close()

	//step2：读取数据
	//bs := make([]byte, 4, 4)
	//将文件拆开几段来读取字节流,对于一些比如中文字符读取容易出现乱码

	bs := make([]byte, 100) //比文件实际需要读取的字节多,会是原本的零值

	n, err := file.Read(bs)
	if err == nil && (n == 0 || err == io.EOF) {
		fmt.Println("读取完毕")
	}

	fmt.Println(err)            //
	fmt.Println(n)              //4
	fmt.Println(bs[:n])         //[97 98 99 100]
	fmt.Println(string(bs[:n])) //abcd
	/*
	    //第一次读取
	   // n,err :=file.Read(bs)
	   // fmt.Println(err) //
	   // fmt.Println(n) //4
	   // fmt.Println(bs) //[97 98 99 100]
	   //fmt.Println(string(bs)) //abcd

	   //第二次读取
	   n,err = file.Read(bs)
	   fmt.Println(err)//
	   fmt.Println(n)//4
	   fmt.Println(bs) //[101 102 103 104]
	   fmt.Println(string(bs)) //efgh

	   //第三次读取
	   n,err = file.Read(bs)
	   fmt.Println(err) //
	   fmt.Println(n) //2
	   fmt.Println(bs) //[105 106 103 104]
	   fmt.Println(string(bs)) //ijgh

	   //第四次读取
	   n,err = file.Read(bs)
	   fmt.Println(err) //EOF
	   fmt.Println(n) //0
	*/
	n = -1
	fmt.Println("-------------")
	for {
		n, err = file.Read(bs) //n是读取的字节数,同一个goroutine中读取文件内容,已经读取完了
		fmt.Println(n)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		//这样比较稳妥,一个字节一个字节标清楚
		fmt.Println(string(bs[:n])) //[]byte->string
	}
}
