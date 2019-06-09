package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 1
	odd := head
	even := head.Next.Next
	evenHead := even
	node := head
	for node != nil {
		if count%2 == 1 {
			odd.Next = node
			odd = node
		} else {
			even.Next = node
			even = node
		}
		node = node.Next
		count++
	}

	odd.Next = evenHead

	return head
}
