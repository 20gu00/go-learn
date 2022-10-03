package main

import "fmt"

func main() {
    /*
       if语句的其他写法：
       if 初始化语句; 条件{
           //注意变量的作用域问题
       }
    */

    //这就是在if语句的判断位置初始化变量，这个变量的作用于是if语句，在走if else整个流程中生效，完成了，变量就销毁了，也就是内存空间收回来了
    //
    //(if else语句的判断条件是有先后顺序的，变量的作用域或者它的作用效果，不能通过在代码中的位置的顺序判断，而是看变量类型来确定作用于)

    if num := 4; num > 0 {
        fmt.Println("正数。。", num)
    } else if num < 0 {
        fmt.Println("负数。。", num)
    }
    //fmt.Println(num) // undefined: num

    num2 := 5
    if num2 > 0 {
        fmt.Println("num2，是正数。。", num2)
    }
    fmt.Println(num2)

}
