package main

import "fmt"

type error_demo struct { //自定义错误类型,其实就是为了包装更多的错误信息
	Msg  string
	File string
}

func (e *error_demo) Error() string {
	return fmt.Sprintf("%s:%s", e.Msg, e.File)
}

func NewErrorDemo() error {
	return &error_demo{
		Msg:  "msg",
		File: "file",
	}
}

func main() {
	err := NewErrorDemo()
	switch err := err.(type) { //右边的err用的是外部变量  类型断言  接口断言的变种switch(type)
	case nil:

	case *error_demo:
		fmt.Println(err)
	default:

	}
}
