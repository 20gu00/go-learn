package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	//结构体指针和结构体往往可以当成一个东西来用,比如访问成员,只是传递时一个值传递一个引用传递
	//结构体调用,不会修改结构体的内容,指针方式调用会修改
	//核心:方法的本质是接受者作为第一个参数的函数,所以一个值传递,深拷贝,数据副本,一个引用传递
	u.ChangeName("Tom Changed!")
	//转换:User.ChangeName(u),注意接受者和参数类型一致,所以就直接是u
	u.ChangeAge(100) //即使接受其实结构体指针,一样也可以用指针调用
	//转换:*User.ChangAge(&u),得是指针,u是结构体,所以加&
	fmt.Printf("%v \n", u)

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!") //即使接收器是结构体,一样可以用结构体指针调用
	//转换:User.ChangeName(*up),这里如果是访问成员的话,直接结构体指针访问也行
	up.ChangeAge(120)
	//转换:*User.ChangeAge(up)

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}
