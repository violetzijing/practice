package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node1 := &TreeNode{
		Val: 1,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  node2,
		Right: node7,
	}
	//	preorderTraversal(node4)
	//	inorderTraversal(node4)
	inorderTraversalLoop(node4)
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println(root.Val)
	inorderTraversal(root.Right)
}

func inorderTraversalLoop(root *TreeNode) {
	node := root
	stack := []*TreeNode{}

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) != 0 {
			node = stack[len(stack)-1]
			fmt.Println(node.Val)
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}

}
