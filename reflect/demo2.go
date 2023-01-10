package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("num的数值:", num)

	//需要操作指针
	//通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	//得是接口或者指针才行，返回他的值
	newValue := pointer.Elem()

	fmt.Println("类型 :", newValue.Type()) //float64
	fmt.Println("类型 :", newValue.Kind())
	fmt.Println("是否可以修改:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("新的数值:", num)

	// 如果reflect.ValueOf的参数不是指针，会如何？

	//尝试直接修改
	//value := reflect.ValueOf(num)
	//value.SetFloat(6.28) //panic: reflect: reflect.Value.SetFloat using unaddressable value
	//fmt.Println(value.CanSet()) //false

	//pointer = reflect.ValueOf(num)
	//newValue = value.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}
