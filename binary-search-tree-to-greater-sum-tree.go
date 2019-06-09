package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	node := root
	stack := []*TreeNode{root}
	sum := 0

	for node != nil || len(stack) == 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}

		if len(stack) != 0 {
			node = stack[len(stack)-1]
			node.Val += sum
			sum = node.Val
			stack = stack[:len(stack)-1]
			node = node.Left
		}
	}
}
