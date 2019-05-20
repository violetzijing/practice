package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node4 := &ListNode{
		Val: 1,
	}
	node3 := &ListNode{
		Val:  2,
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
	fmt.Println(isPalindrome2(node1))
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	stack := []int{}
	node := head
	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	fmt.Println("stack: ", stack)

	i := 0
	j := len(stack) - 1
	for i < j {
		if stack[i] != stack[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}
	slow = reverse(slow)
	fast = head
	for slow != nil {
		fmt.Println("slow.Val: ", slow.Val)
		fmt.Println("fast.Val: ", fast.Val)
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
