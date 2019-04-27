package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node4 := &ListNode{
		Val: 4,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	fmt.Println(middleNode(node1))
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 0

	node := head
	for node != nil {
		length++
		node = node.Next
	}
	mid := (length / 2)
	node = head
	for mid != 0 {
		node = node.Next
		mid--
	}

	return node
}
