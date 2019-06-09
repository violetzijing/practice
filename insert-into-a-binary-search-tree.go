package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	node := root
	for {
		if node.Val <= val {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{Val: val}
				break
			}
		}
	}

	return root
}
