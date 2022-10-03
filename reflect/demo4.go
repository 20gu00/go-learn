package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age int
    Sex string
}

func (p Person)Say(msg string)  {
    fmt.Println("hello，",msg)
}
func (p Person)PrintInfo()  {
    fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}

func (p Person) Test(i,j int,s string){
    fmt.Println(i,j,s)
}

// 如何通过反射来进行方法的调用？
// 本来可以用结构体对象.方法名称()直接调用的，
// 但是如果要通过反射，
// 那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
    p2 := Person{"Ruby",30,"男"}
    // 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，
    // 得到“反射类型对象”后才能做下一步处理
    getValue := reflect.ValueOf(p2)

    // 2.一定要指定参数为正确的方法名
    // 先看看没有参数的调用方法

    methodValue1 := getValue.MethodByName("PrintInfo")
    fmt.Printf("Kind : %s, Type : %s\n",methodValue1.Kind(),methodValue1.Type())
    methodValue1.Call(nil) //没有参数，直接写nil

    args1 := make([]reflect.Value, 0) //或者创建一个空的切片也可以
    methodValue1.Call(args1)

    // 有参数的方法调用
    methodValue2 := getValue.MethodByName("Say")
    fmt.Printf("Kind : %s, Type : %s\n",methodValue2.Kind(),methodValue2.Type())
    args2 := []reflect.Value{reflect.ValueOf("反射机制")}
    methodValue2.Call(args2)

    methodValue3 := getValue.MethodByName("Test")
    fmt.Printf("Kind : %s, Type : %s\n",methodValue3.Kind(),methodValue3.Type())
    args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200),reflect.ValueOf("Hello")}

    methodValue3.Call(args3)
}

