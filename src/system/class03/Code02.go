package class03

import "fmt"

type DoubleNode struct {
	Value int
	Next  *DoubleNode
	Pre   *DoubleNode
}

func PrintLink2(node *DoubleNode) {

	for true {
		if node == nil {
			break
		}
		fmt.Println(node.Value)
		node = node.Next
	}
}

func DelDoubleNode(head *DoubleNode, val int) *DoubleNode {

	if nil == head {
		return nil
	}

	if head.Value == val {
		head.Next.Pre = nil
		return head.Next
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Value == val {
			cur.Next = cur.Next.Next
			cur.Next.Pre = cur
		}
		cur = cur.Next
	}

	return head
}

func ReverseDoubleNode(head *DoubleNode) *DoubleNode {
	if nil == head {
		return nil
	}

	var pre *DoubleNode
	var cur = head
	var next = head.Next
	for next != nil {
		cur.Next = pre
		cur.Pre = next
		pre = cur
		cur = next
		next = next.Next
	}

	cur.Next = pre
	cur.Pre = nil

	return cur
}
