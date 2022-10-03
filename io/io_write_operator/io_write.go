package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
	   写出数据：
	*/

	fileName := "/root/go-learn/io/io_write_operator/test.txt"
	//step1：打开文件
	//step2：写出数据
	//step3：关闭文件
	//file,err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm) //os.O_CREATE|os.O_WRONLY|os.O_APPEND
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//写出数据
	//bs :=[]byte{65,66,67,68,69,70}//A,B,C,D,E,F

	//或者这里你可以先定义个字符串,在string->[]byte
	bs := []byte{97, 98, 99, 100} //a,b,c,d  字符串,go采用unicode字符集和utf-8编码,这里包含了ASCII
	//n,err := file.Write(bs)
	n, err := file.Write(bs[:2]) //写入的是字节流
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n") //写入个换行符,之所以叫写出数据,是因为一切都是io的字节流槽错

	//直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)

	//如果开启了append,则seek设置无效了其实,要注意下,这个其实覆盖掉已有的字节流
	//file.Seek(5, 0)
	//n, err = file.Write([]byte("seek"))
	n, err = file.WriteAt([]byte("seekat"), 1) //就是第一个字节流后,包括换行符也是字节
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
