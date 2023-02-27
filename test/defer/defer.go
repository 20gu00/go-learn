package main

import "fmt"

func main() {
	test("jo")
}

func test(name string) {
	defer func(p string) {
		fmt.Println(p)
	}(name)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	name = "Lee"
	panic("error")
	//下面的代码块不可以达到的了
	defer func() {
		fmt.Println("end")
	}()
}
