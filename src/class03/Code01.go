package class03

import "fmt"

// Node 单链表节点
type Node struct {
	Value int
	Next  *Node
}

// LinkedList 单链表
type LinkedList struct {
	Size int
	Head *Node
}

func AddNode(list *LinkedList, val int) {

	head := new(Node)
	head.Value = val
	list.Size++

	if list.Head == nil {
		list.Head = head
		return
	}

	cur := list.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
}

// PrintLink 打印单链表
func PrintLink(node *Node) {

	for true {
		if node == nil {
			break
		}
		fmt.Println(node.Value)
		node = node.Next
	}
}

// DelNode 单链表删除节点
func DelNode(list *LinkedList, val int) {

	head := list.Head

	if nil == list {
		return
	}

	if head.Value == val {
		list.Size--
		list.Head = list.Head.Next
		return
	}

	pre := head
	cur := head
	for cur != nil {
		if cur.Value == val {
			list.Size--
			pre.Next = cur.Next
		}
		pre = cur
		cur = cur.Next
	}

}

// ReverseLinked 单链表反转节点
func ReverseLinked(head *Node) *Node {
	if nil == head {
		return nil
	}

	var pre *Node
	var next *Node

	for nil != head.Next {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	head.Next = pre

	return head

}

func ReverseLinked2(list *LinkedList) {
	if nil == list.Head {
		return
	}

	var pre *Node
	var mid *Node
	var next *Node

	mid = list.Head
	for nil != mid.Next {
		next = mid.Next
		mid.Next = pre
		pre = mid
		mid = next
	}
	mid.Next = pre
	list.Head = mid

}
