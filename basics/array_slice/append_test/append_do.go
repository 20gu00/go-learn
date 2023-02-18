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
	return
}

func m1(t []int) []int {
	t[1] = 100
	return t
}