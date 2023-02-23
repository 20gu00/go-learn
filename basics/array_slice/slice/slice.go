package main

import "fmt"

func main() {
	/*
		数组array：
			存储一组相同数据类型的数据结构。
				特点：定长

		切片slice：
			同数组类似，也叫做变长数组或者动态数组。
				特点：变长

			是一个引用类型的容器，指向了一个底层数组。

		make()
			func make(t Type, size ...IntegerType) Type

			第一个参数：类型
				slice，map，chan
			第二个参数：长度len
				实际存储元素的数量
			第三个参数：容量cap
				最多能够存储的元素的数量


		append()，专门用于向切片的尾部追加元素
			slice = append(slice, elem1, elem2)
			slice = append(slice, anotherSlice...)
	*/
	//数组一旦定义了或者说赋值了(因为有时候数组时还不指定长度)，那么数组长度就固定了
	//而切片，你可以把它理解成动态数组
	//切片是一个引用类型的容器，指向底层的数组
	//
	//对于切片，建议使用make来创造，make专门用来创造引用类型的数据
	//
	//数组是定长的，所以len()和cap()一样

	//1.数组
	arr := [4]int{1, 2, 3, 4} //定长
	fmt.Println(arr)

	//2.切片
	var s1 []int //这样定义切片，默认是空
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4} //变长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2) //[4]int,[]int
	//切片的删除
	s2 = s2[1:] //删除切片的第一个值
	fmt.Println("s2:", s2)

	s3 := make([]int, 3, 8) //这样定义切片，默认值是[0,0,0]
	fmt.Println(s3)
	fmt.Printf("容量：%d,长度：%d\n", cap(s3), len(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)
	//fmt.Println(s3[3]) //panic: runtime error: index out of range

	//append()//向切片末尾添加元素
	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1, 2) //append操作切片，可能会使切片扩容，所以要让切片s4重新指向
	fmt.Println(s4)
	s4 = append(s4, 3, 4, 5, 6, 7) // append(s4,[1,2,3,4,5]...)
	fmt.Println(s4)

	s4 = append(s4, s3...) //s3...表示这不是s3而是s3的元素(注意元素类型)
	fmt.Println(s4)

	//遍历切片
	for i := 0; i < len(s4); i++ { //注意这里是长度
		fmt.Println(s4[i])
	}

	for i, v := range s4 {
		fmt.Printf("%d-->%d\n", i, v)
	}

}
