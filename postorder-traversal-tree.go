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

	//	postorderTraversal(node4)
	postorderTraversalLoop(node4)
}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println(root.Val)
}

func postorderTraversalLoop(root *TreeNode) {
	first := []*TreeNode{}
	second := []*TreeNode{}
	node := root
	first = append(first, node)

	for node != nil && len(first) != 0 {
		node = first[len(first)-1]
		first = first[:len(first)-1]
		second = append(second, node)
		if node.Left != nil {
			first = append(first, node.Left)
		}
		if node.Right != nil {
			first = append(first, node.Right)
		}
	}

	for _, item := range second {
		fmt.Println(item.Val)
	}
}
