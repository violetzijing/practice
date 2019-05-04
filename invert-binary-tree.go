package main

import "fmt"

func main() {
	node1 := &TreeNode{
		Val: 1,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node6 := &TreeNode{
		Val: 6,
	}
	node9 := &TreeNode{
		Val: 9,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}
	node7 := &TreeNode{
		Val:   7,
		Left:  node6,
		Right: node9,
	}
	node4 := &TreeNode{
		Val:   4,
		Left:  node2,
		Right: node7,
	}
	fmt.Println(invertTree(node4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		invertTree(root.Left)
		invertTree(root.Right)
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
	}

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
			tmp := node.Left
			node.Left = node.Right
			node.Right = tmp
		}
	}

	return root
}
