package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error = &MyError{}
	println(err.Error())

	ErrorsPkg()
}

//其实很少自己实现各错误类型
type MyError struct {
}

//自定义的错误类型结构体实现了Error方法,自己定义的同名的方法
func (m *MyError) Error() string {
	return "Hello, it's my error"
}

func ErrorsPkg() {
	err := &MyError{}
	// 使用 %w 占位符，返回的是一个新错误  返回个包装
	// wrappedErr 是一个新类型，fmt.wrapError
	//fmt.Printxxx 输出到控制台
	//fmt.Sprintxxx 返回string
	//fmt.Errorxxx 返回error
	wrappedErr := fmt.Errorf("this is an wrapped error %w", err)

	// 再解出来,可以重复调用,调用一次解一次包,知道最原始的错误
	if err == errors.Unwrap(wrappedErr) {
		fmt.Println("unwrapped") //未包装,解包
	}

	if errors.Is(wrappedErr, err) { //用Is判断是不是相等,注意这里会自己层层解包
		// 虽然被包了一下，但是 Is 会逐层解除包装，判断是不是该错误
		fmt.Println("wrapped is err")
	}

	copyErr := &MyError{}
	// 这里尝试将 wrappedErr转换为 MyError
	// 注意我们使用了两次的取地址符号
	if errors.As(wrappedErr, &copyErr) {
		fmt.Println("convert error")
	}
}
