package main

import "fmt"

func main() {
	n := []int{1, 9, 3, 6, 0, 10}
	tar := 10
	out := func(nums []int, target int) []int {
		//空的切片
		//r := []int{}  不建议用这种方式创建切片
		r := make([]int, 0, 10)
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if t := nums[i] + nums[j]; t == target {
					//r := []int{j, j + 1}
					r = append(r, i, j)
				}

			}
		}
		return r
	}(n, tar)
	fmt.Println(out)
}
