package main

import "fmt"

func main() {
	node9 := &ListNode{
		Val: 9,
	}
	node6 := &ListNode{
		Val:  6,
		Next: node9,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node6,
	}

	node5 := &ListNode{
		Val:  5,
		Next: node1,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node4,
	}

	fmt.Println("Before:")
	Print(node2)
	SortList(node2)
	fmt.Println("After:")
	fmt.Println(Print(node2))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head.Next == nil {
		// Only 1 element. Already sorted
		return head
	}
	fast := head
	slow := head
	tail := slow // The tail of the first part
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		tail = slow
		slow = slow.Next
	}
	// First part
	a := slow
	tail.Next = nil
	b := head
	leftStr := Print(b)
	rightStr := Print(a)
	a = SortList(a)
	b = SortList(b)
	head = Merge(a, b)
	mergedStr := Print(head)
	fmt.Printf("(%s) + (%s) => (%s)\n", leftStr, rightStr, mergedStr)
	return head
}

func Merge(a, b *ListNode) *ListNode {
	var (
		head *ListNode
		p1   *ListNode
		p2   *ListNode
	)

	if a.Val > b.Val {
		head = b
		p1 = a
		p2 = b.Next
	} else {
		head = a
		p1 = a.Next
		p2 = b
	}

	node := head
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			node.Next = p2
			p2 = p2.Next
		} else {
			node.Next = p1
			p1 = p1.Next
		}
		node = node.Next
	}

	if p1 != nil {
		node.Next = p1
	}
	if p2 != nil {
		node.Next = p2
	}

	Print(head)

	return head
}

func Print(head *ListNode) string {
	str := ""
	for head != nil {
		str = fmt.Sprintf("%s -> %d", str, head.Val)
		head = head.Next
	}
	return str
}
