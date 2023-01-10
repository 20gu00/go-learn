package main

import (
	"fmt"
	"reflect"
)

func main() {
	//reflect运行期间获取对象的类型信息和内存结构
	//反射操作：通过反射，可以获取一个接口类型变量的 类型和数值
	//reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
	//    reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 “Allen.Wu” 25} 这样的结构体struct的值
	//    也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种
	//    Type 和 Value 都包含了大量的方法，其中第一个有用的方法应该是 Kind，这个方法返回该类型的具体信息：Uint、Float64 等。Value 类型还包含了一系列类型方法，比如 Int()，用于返回对应的值。
	var x float64 = 3.4

	fmt.Println("type:", reflect.TypeOf(x))   //type: float64
	fmt.Println("value:", reflect.ValueOf(x)) //value: 3.4

	fmt.Println("-------------------")
	//根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("type : ", v.Type())
	fmt.Println("value : ", v.Float())
}
