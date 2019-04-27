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
	find := 5
	fmt.Println("result: ", searchBST(node4, find))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	node := root
	for node != nil {
		if node.Val == val {
			return node
		}
		if node.Val > val {
			node = node.Left
			continue
		}
		if node.Val < val {
			node = node.Right
		}
	}

	return nil
}
