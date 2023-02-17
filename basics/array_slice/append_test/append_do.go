package main

import "fmt"

func main() {
	s := []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	q := append(s, 4) //实际上就是s是[1,2,3],然后这里是新的切片,那么就是将s拷贝(深拷贝)值给q,然后对q在len(s)长度的地方也就是的3索引开始添加新的元素
	w := append(s, 5)
	r := append(q, 11)
	fmt.Println("s:", s)
	fmt.Println("q:", q)
	fmt.Println("w:", w)
	fmt.Println("r:", r)

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
