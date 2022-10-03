package main

import (
    "reflect"
    "fmt"
)

type Student struct {
    Name string
    Age int
    School string
}
func main()  {
    /*
    通过反射，来更改对象的数值：前提是数据可以被更改
     */
    s1:=Student{"王二狗",19,"千锋教育"}
    fmt.Printf("%T\n",s1) //main.Student
    p1:=&s1
    fmt.Printf("%T\n",p1) //*main.Student
    fmt.Println(s1.Name)
    fmt.Println((*p1).Name,p1.Name)

    v1:= reflect.ValueOf(&s1) // value

    if v1.Kind()==reflect.Ptr{
        fmt.Println(v1.Elem().CanSet())
        v1 = v1.Elem()
    }

    f1:=v1.FieldByName("Name")
    f1.SetString("韩茹")
    f3:=v1.FieldByName("School")
    f3.SetString("幼儿园")
    fmt.Println(s1)

}

