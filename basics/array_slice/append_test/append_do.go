package main

import (
	"fmt"
	"time"
)

func main() {
	// 留意超出容量会更换底层数据,但是append创建切片中,整体上看,依旧是共享数组
	// append的参数可以是一个切片 ... 批量插入
	s := []int{}
	s = append(s, 1)
	fmt.Println("---", len(s), cap(s))
	s = append(s, 2)
	fmt.Println("---", len(s), cap(s))
	s = append(s, 3)
	fmt.Println("---", len(s), cap(s))
	q := append(s, 4) //实际上就是s是[1,2,3],然后这里是新的切片,那么就是将s拷贝(深拷贝)值给q,然后对q在len(s)长度的地方也就是的3索引开始添加新的元素
	fmt.Println("---", len(s), cap(s))
	fmt.Println("---", len(q), cap(q))
	w := append(s, 5) // 同一个数组,修改那个位置,覆盖了元素
	fmt.Println("1---", len(s), cap(s))
	fmt.Println("1---", len(q), cap(q))
	fmt.Println("1---", len(w), cap(w))
	r := append(q, 11) //s,q,w同一个底层数组,r不是
	fmt.Println("2---", len(s), cap(s))
	fmt.Println("2---", len(q), cap(q))
	fmt.Println("2---", len(w), cap(w))
	fmt.Println("3---", len(r), cap(r))

	fmt.Println("s:", s)
	fmt.Println("q:", q)
	fmt.Println("w:", w)
	fmt.Println("r:", r)

	s[0] = 1000
	fmt.Println(s, q, w, r) //[1000 2 3] [1000 2 3 5] [1000 2 3 5] [1 2 3 5 11]

	// 数组地址实际上就是第一个元素地址
	// 这三个切片都指向了同一个数组,只是len有所不同,实际上切片就是struct,由pointer len cap
	//s: [1 2 3]
	//q: [1 2 3 5]
	//w: [1 2 3 5]
	//r: [1 2 3 5 11]

	/*
			append每次返回的是一个新的切片,看你用新的还是就得切片变量来接收
			上面的s切片,一开始没有设置容量,添加1进去,注意底层数组长度永远是固定,长度变为1,容量变为1.
		    对s添加2进去,长度变为2,容量变为2,底层数组变更了
			对s添加3进去,扩容,底层数组变更,容量为4,切片长度为3
			添加4进去,但是使用的是新的切片变量q,长度为4,容量为超出,所以s和q同一个底层数组,s[1,2,3],q[1,2,3,4]
			添加5进去,但是使用的是新的切片变量w,长度为4,容量未超出,所以w和q同一个底层数组,而且修改了数组索引3也就是第四个元素s[1,2,3],q[1,2,3,5],w[1,2,3,5]
			对q添加11进去,长度为5,扩容,新的数组,所以r和sqw不同数组

			append只需要注意三点:
			1.扩容了吗
			2.每次都创建一个新的切片,就看是用新旧切片变量来接收
			3.实际上是对第一个参数进行一次值拷贝
	*/

	//append删除元素可以通过索引实现,比如长度为9[1,2,3,4,5,6,7,8,9]
	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t = append(t[0:3], t[5:]...) //[1 2 3 6 7 8 9]
	fmt.Println(t)

	t3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t2 := append(t3[0:3], t3[5:]...) //[1 2 3 6 7 8 9]
	fmt.Println(t2, t3)              //[1 2 3 6 7 8 9] [1 2 3 6 7 8 9 8 9] 底层数组没有改变,为扩容
	//----------------------------------------------------------------------------

	//s2 := make([]int, 10)  // 这是有十个零值,append往后面加
	s2 := make([]int, 0, 10)
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	q2 := append(s2, 4)
	w2 := append(s2, 5)
	r2 := append(q2, 11)
	fmt.Println("s2:", s2)
	fmt.Println("q2:", q2)
	fmt.Println("w2:", w2)
	fmt.Println("r2:", r2)

	s2[0] = 1111
	fmt.Println(s2, q2, w2, r2) //[1111 2 3] [1111 2 3 5] [1111 2 3 5] [1111 2 3 5 11] 全都同一个底层数组

	//s2: [1 2 3]
	//q2: [1 2 3 5]
	//w2: [1 2 3 5]
	//r2: [1 2 3 5 11]

	// 引用类型
	s3 := []int{1, 2}
	fmt.Println("s3:", s3)
	s4 := m1(s3)
	fmt.Println("s4:", s4)
	fmt.Println("s3:", s3)

	// for range
	fruits := []string{"a", "b", "c", "d", "e"}
	for _, v := range fruits {
		go func() {
			fmt.Println(v) // e e e e e  因为for循环很可能已经执行已经执行完了,v是e,协程才打印
		}()
		//time.Sleep(1 * time.Second) a b c d e 阻塞一会,让协程运行完再继续for循环(不像defer,协程没有提前运算也就是副本)
	}
	time.Sleep(10 * time.Second)
	//select{}

	aq := make([]int, 0, 20)
	aq = append(aq, 1)
	aq = append(aq, 2)
	//am2 := m2(aq)
	m2(aq)
	//这是在底层数据添加了一个300,实际上覆盖了原有3,aq和m2的a这里都是300
	aq = append(aq, 300)
	//fmt.Println(am2)
	fmt.Println(aq)

	return
}

func m1(t []int) []int {
	t[1] = 100
	return t
}

//一句话总结：不管是值类型还是引用类型,除非使用指针,Go语言中的函数传参方式全部都是值传递，不存在引用传递。
//函数里面通过索引号修改可以修改外部函数也就是原切片即底层数组的值,但是删除或者插入修改的是函数里面的切片而不是外部,也就是不影响原本的底层数组,环底层数组了
//实际上就是struct拷贝,两个切片了
//可以传递个指针,或者返回一个切片,不然就只有通过索引会影响外部的切片,所以外部的切片这时候不受影响len cap都不变,除非外部自己更改
func m2(a []int) []int {
	fmt.Println(len(a), cap(a)) //2 20
	a = append(a, 3)
	a = append(a, 4)
	a[0] = 100
	fmt.Println(len(a), cap(a)) // 4 20
	fmt.Println("a", a)
	return a
}

/*
	方法的值接受者和指针接受者
	方法的本质是接受者作为第一个参数

	//下面的对象都是结构体类型或者结构体针织类型,而不是直接interface类型
	值对象(var z z=z{})可以调用值接受者的方法甚至指针接受者(&z)的方法
	指针对象(z:=&z{})可以调用值接受者(*z)的方法和指针接受者的方法

	如果有个接口,里边的方法有某个结构体都实现了,有值接受者也有指针接受者:
	type i interface{
		a()
		b()
	}

	type s struct{}

	func (s s)a(){
		fmt.Println("a")
	}

	func (s *s)b(){
		fmt.Println("b")
	}

	main中:
	// 直接实现接口类型而不是结构体类型
	var s2 i=&s{}
	//s2可以调用的a() b()
	s2.a()
	s2.b()



	//如果是值类型实现的接口
	var s3 i=s{}
	s3.a()
	s3.b()  //不能调用b()方法,因为接受者是指针类型
	// 编译会报错
	//建议统一使用指针类型
*/
