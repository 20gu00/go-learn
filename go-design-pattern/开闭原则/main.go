package main

import "fmt"

//开闭原则
//银行业务员banker,比如说功能有各种业务 支付 转账 现在又加一个股票
//比如我有一个类,它有三个方法,也就是三个功能,那如果我要添加一个功能,那么就是添加一个方法,但是这样修改了类的代码
//尤其是类很庞大的时候,你添加个方法,很可能会影响其他方法

//开闭原则,比如将业务banker进行抽象成一个接口,根据这个抽象模去实现支付banker或者转账banker,类似继承,添加一个股票banker也是这样,但是不会修改稳定的代码,不会影响其他模块比如支付转账等等
//这就是迭代去开发,不修改原本的代码
//就是抽象,面向接口编程,不修改原本的代码,而是重写接口

//但是接口中的方法一致,也就是参数和返回值

//抽象的banker
type AbBanker interface {
	//方法 功能
	DoBusy() // 业务
}

//interface是抽象,struct是具体的

//存款功能,也就是有存储功能的banker
//实现这个接口,实现这个抽象,面向接口编程
type SaveBanker struct {
	//AbBanker
}

//实现接口的方法,可以理解成重写这个接口
func (*SaveBanker) DoBusy() {
	fmt.Println("存款")
}

func main() {
	//调用各种功能
	//这里最好复制给接口类型的变量
	s := SaveBanker{}
	s.DoBusy()

	var t AbBanker
	t = &TranBanker{}
	t.DoBusy()

	Busy(&s)
	Busy(t)
}

//添加转账功能,直接添加一个新的类,重写这个接口
//额外添加一个功能,但是不修改原本的代码,或者已有的类
type TranBanker struct {
	//AbBanker
}

func (*TranBanker) DoBusy() {
	fmt.Println("转账")
}

//还可以实现一个架构层
func Busy(banker AbBanker) {
	banker.DoBusy()
}
