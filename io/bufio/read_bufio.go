package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	   bufio:高效io读写
	       buffer缓存
	       io：input/output

	   将io包下的Reader，Write对象进行包装，带缓存的包装，提高读写的效率

	       ReadBytes()
	       ReadString()
	       ReadLine()
	*/
	//注意,打开文件进行读取写入,都是不可重复的,意思是比如你有代码块读取了文件内容出来
	//你下一个代码块读取文件是空的,err会显示EOF,读取到的内容一般是零值,比如你到读取字符串readString方法,那就是空的字符串

	fileName := "/root/go-learn/io/bufio/test"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//创建Reader对象
	b1 := bufio.NewReader(file)
	//1.Read()，高效读取
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))

	//2.ReadLine()
	data, flag, err := b1.ReadLine() //按行读取,不太好用
	fmt.Println("readline")
	fmt.Println(flag)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	//3.ReadString()
	s1, err := b1.ReadString('\n') //分隔符
	fmt.Println(err)
	fmt.Println(s1)

	// s1,err = b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)

	//s1,err = b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	for {
		s1, err := b1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕。。")
			break
		}
		fmt.Println(s1)
	}

	//4.ReadBytes()
	data, err = b1.ReadBytes('\n')
	fmt.Println(err)
	fmt.Println(string(data))

	//Scanner
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)

	//从标准输入读取os.Stdin
	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)

}
