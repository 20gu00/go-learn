package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()
	//上面这种方式定义的指针没问题,这种情况指针是有值的,一般是var这种定义指针便来那个的方式没有赋值

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	duck3 := new(ToyDuck)
	duck3.Swim()

	// 当你声明这样的时候，Go 就帮你分配好内存
	// 不用担心空指针的问题，以为它压根就不是指针
	//其实就是bit位全是0,对应类型零值
	var duck4 ToyDuck
	duck4.Swim()

	//分配了内存但是指针没有指向任何内存地址,所以就是空指针,严谨点说是nil
	//声明了一个指针但是没赋值,那么就是nil
	// duck5 就是一个指针了
	var duck5 *ToyDuck
	// 这边会直接panic 掉
	duck5.Swim()

	// 赋值，初始化按字段名字赋值
	duck6 := ToyDuck{
		Color: "黄色",
		Price: 100,
	}
	duck6.Swim()

	// 初始化按字段顺序赋值，不建议使用,结构体一改这里也需要改
	duck7 := ToyDuck{"蓝色", 1024}
	duck7.Swim()

	// 后面再单独赋值
	duck8 := ToyDuck{}
	duck8.Color = "橘色"

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s，%d一只\n", t.Color, t.Price)
}
