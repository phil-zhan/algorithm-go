package main

import (
	"algorithm-go/src/class03"
	"fmt"
)

func main() {
	list := new(class03.LinkedList)
	for i := 0; i < 10; i++ {
		class03.AddNode(list, i)
	}
	class03.DelNode(list, 3)
	class03.PrintLink(list.Head)

	println("=================================")

	class03.ReverseLinked2(list)
	class03.PrintLink(list.Head)

	println("===============================下面是双链表===============================")
	// 后面是双节点
	doubleNode1 := new(class03.DoubleNode)
	doubleNode1.Value = 1
	doubleNode2 := new(class03.DoubleNode)
	doubleNode2.Value = 2
	doubleNode3 := new(class03.DoubleNode)
	doubleNode3.Value = 3
	doubleNode4 := new(class03.DoubleNode)
	doubleNode4.Value = 4
	doubleNode5 := new(class03.DoubleNode)
	doubleNode5.Value = 5
	doubleNode6 := new(class03.DoubleNode)
	doubleNode6.Value = 6
	doubleNode7 := new(class03.DoubleNode)
	doubleNode7.Value = 7

	doubleNode1.Next = doubleNode2
	doubleNode2.Next = doubleNode3
	doubleNode3.Next = doubleNode4
	doubleNode4.Next = doubleNode5
	doubleNode5.Next = doubleNode6
	doubleNode6.Next = doubleNode7

	doubleNode2.Pre = doubleNode1
	doubleNode3.Pre = doubleNode2
	doubleNode4.Pre = doubleNode3
	doubleNode5.Pre = doubleNode4
	doubleNode6.Pre = doubleNode5
	doubleNode7.Pre = doubleNode6

	println("删除前")
	class03.PrintLink2(doubleNode1)
	println("删除前")

	doubleNodeDel := class03.DelDoubleNode(doubleNode1, 3)

	println("删除后")
	class03.PrintLink2(doubleNodeDel)
	println("删除后")

	doubleNodeRev := class03.ReverseDoubleNode(doubleNode1)
	class03.PrintLink2(doubleNodeRev)

	println("===============================下面是栈结构===============================")
	stack := new(class03.Stack)
	for i := 0; i < 10; i++ {
		class03.PutStack(stack, i)
	}

	size := class03.StackSize(*stack)
	for i := 0; i < size; i++ {
		println(class03.PopStack(stack))
	}

	println("===============================下面是递归===============================")
	// 测试递归最大值
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	max := class03.GetMax(arr)
	fmt.Println(max)
}
