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
	preorderTraversalLoop(node4)
}

func preorderTraversal(root *TreeNode) {
	fmt.Println(root.Val)
	if root.Left == nil || root.Right == nil {
		return
	}
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func preorderTraversalLoop(root *TreeNode) {
	list := []*TreeNode{}
	node := root
	for node != nil || len(list) != 0 {
		for node != nil {
			fmt.Println(node.Val)
			list = append(list, node)
			node = node.Left
		}
		if node == nil && len(list) != 0 {
			node = list[len(list)-1]
			list = list[:len(list)-1]
			node = node.Right
		}
	}
}
