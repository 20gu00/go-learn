package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map的遍历：
			使用：for range

				数组，切片：index，value
				map：key，value

		只能用for_range遍历map,同时具有随机性
	*/

	//go中使用sort包进行排序很方便,想原始就自己使用两个for循环进行排序
	map1 := make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "金角大王"
	map1[6] = "王二狗"

	//1.遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("----------------------")
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, "--->", map1[i])
	}
	/*
		1.获取所有的key，-->切片/数组
		2.进行排序
		3.遍历key，--->map[key]
	*/
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//冒泡排序，或者使用sort包下的排序方法
	//sort.Ints()操作对象是slice,从小到大排列
	sort.Ints(keys)
	fmt.Println("sort:", keys)
	//reverse
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Println("sort by desc:", keys)
	//use for 注意这里比较的是值,在我们这里也中文字符,只是进行效果演示
	//for i := 1; i <= 6; i++ {
	//	for j := 0; j < len(keys)-1; i++ {
	//		if keys[j] < keys[j+1] {
	//			keys[j], keys[j+1] = keys[j+1], keys[j]
	//		}
	//	}
	//}
	//fmt.Println("use for", keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"Apple", "Windows", "Orange", "abc", "王二狗", "acd", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
