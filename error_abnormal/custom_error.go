package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		自定义错误：
	*/
	radius := -3.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok { //接口的断言,接口的动态类型
			fmt.Printf("半径是：%.2f\n", err.radius)
		}
		return
	}
	fmt.Println("圆形的面积是：", area)

}

//1.定义一个结构体，表示错误的类型
type areaError struct {
	msg    string
	radius float64
}

//2.实现error接口，就是实现Error()方法
//因为go的内置的error类型是接口类型,要求实现Error() string方法,这里我们也实现它即可
//也实现这个方法,那也就是我们定义的结构体对象也实现了error接口

//有的类型比如基础类型int这些他不能作为方法的接受者,所以int一般不能实现接口,也就是不能作为接口的断言的对象,除非是interface{},因为任何数据类型都能实现interface{}
func (e *areaError) Error() string { //这里是结构体指针,实现接口,所以断言也是断言结构体指针
	return fmt.Sprintf("error：半径，%.2f，%s", e.radius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法的", radius} //&使用结构体
	}
	return math.Pi * radius * radius, nil
}
