package main

import "fmt"

func main() {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node1,
	}
	node1.Next = node2
	fmt.Println(hasCycle(node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}

	}

	return false
}
