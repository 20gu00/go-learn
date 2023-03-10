package main

import "fmt"

//单一职责原则
//平时可能一个结构体有两个方法,但是有可能两个方法内部内容实际上是一样的,阅读者就容易混乱,本来是两个业务也就是对应两个不同方法,现在看来不同的业务统一方法也可以搞定了
//比如一个类也就是一个struct clothes有两个方法Onwork和Onshop,业务就是穿衣服不一样

//现在用单一原则来实现,一个结构体,就一个方法,也就是clothesWork对应一个Onwork方法,一个clothesShop对应一个Onshop方法
//两个结构体两个方法,每个类只有一个方法,,对外提供一种功能,单一职责

type CloWork struct{}

func (cw *CloWork) Onwork() {
	fmt.Println("onwork穿着")
}

type CloShop struct{}

func (cs *CloShop) OnShop() {
	fmt.Println("onshop穿着")
}

func main() {
	//work业务
	cw := CloWork{}
	cw.Onwork()

	//shop业务
	cs := CloShop{}
	cs.OnShop()
}
