package main

import "fmt"

func main() {
	// 留意超出容量会更换底层数据,但是append创建切片中,整体上看,依旧是共享数组
	// append的参数可以是一个切片 ... 批量插入
	s := []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	q := append(s, 4)  //实际上就是s是[1,2,3],然后这里是新的切片,那么就是将s拷贝(深拷贝)值给q,然后对q在len(s)长度的地方也就是的3索引开始添加新的元素
	w := append(s, 5)  // 同一个数组,修改那个位置,覆盖了元素
	r := append(q, 11) //s,q,w同一个底层数组,r不是
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
	return
}

func m1(t []int) []int {
	t[1] = 100
	return t
}
