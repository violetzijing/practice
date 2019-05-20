package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		a = a.Next
		b = b.Next
		if a.Val == b.Val {
			return a
		}
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
	}

	return nil
}
