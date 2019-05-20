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

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow2 := head
			for slow2 != slow {
				slow2 = slow2.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
