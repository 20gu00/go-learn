package main

func main() {

}

type Node struct {
	//自引用只能使用指针
	//结构体不能直接自己引用自己,也不能互相引用,即a引用b,b也引用a,因为编译时要给这个结构体分配内存大小,但这样不能确定内存大小,但指针类型的大小是知道的
	//left Node
	//right Node

	left  *Node
	right *Node

	// 这个也会报错
	// nn NodeNode
}

type NodeNode struct {
	node Node
}
