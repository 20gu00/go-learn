package main

import (
    "fmt"
    "reflect"
)

func main() {
    //函数的反射
    f1 := fun1
    value := reflect.ValueOf(f1)
    fmt.Printf("Kind : %s , Type : %s\n",value.Kind(),value.Type()) //Kind : func , Type : func()

    value2 := reflect.ValueOf(fun2)
    fmt.Printf("Kind : %s , Type : %s\n",value2.Kind(),value2.Type()) //Kind : func , Type : func(int, string)

    //通过反射调用函数
    value.Call(nil)

    value2.Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf("hello")})

}

func fun1(){
    fmt.Println("我是函数fun1()，无参的。。")
}

func fun2(i int, s string){
    fmt.Println("我是函数fun2()，有参数。。",i,s)
}

