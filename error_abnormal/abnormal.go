package main

import (
	"fmt"
	"os"
)

//如果出现错误会报错同时中断代码的运行
//错误是意料之中，就是在可能出现问题的地方出现的问题，比如打开一个文件时打开文件失败，就是一种错误
//
//异常是意料之外，在不应该出现问题的地方出现了问题，比如引用了空指针
//
//错误是业务流程的一部分，而异常不是，错误是一种类型，用内置error类型表示，错误值可以存储在变量中，从函数中返回等等，就和其他的类型一样
//os包这些的函数一般都会有返回个错误服务

//如果一个函数或者方法返回一个错误，按照惯例，它必须是函数返回的最后一个值
//处理错误的管用方法是将返回的错误和nil进行比较，nil表示没有错误，非nil表示出现错误

func main() {
	//test.txt  ./test.txt  /root/go-learn/error_abnormal/test.txt
	f, err := os.Open("test.txt") //想看这个包的这个函数，Ctrl键点击Open
	if err != nil {
		//log.Fatal(err)  //记录错误,然后直接退出程序,defer也不执行
		fmt.Println(err)                        //open test.txt: no such file or directory，这只是简单打印了错误的描述
		if ins, ok := err.(*os.PathError); ok { //接口的断言,这里是个常量,该常量隐式定义,即类型推断,值是个结构体,这里其实就是断言是不是该结构体指针
			fmt.Println("1.Op:", ins.Op)     //操作
			fmt.Println("2.Path:", ins.Path) //路径
			fmt.Println("3.Err:", ins.Err)   //错误描述
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")
	//f为文件的对象,f.name()方法可以获取文件的名称,主要是路径名,即打开是用的路径
}
