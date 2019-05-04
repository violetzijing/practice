package main

import "fmt"

func main() {
	node15 := &TreeNode{
		Val: 15,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node20 := &TreeNode{
		Val:   20,
		Left:  node15,
		Right: node7,
	}
	node9 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  node9,
		Right: node20,
	}
	fmt.Println(levelOrderBottom(node3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{
		root,
	}
	result := [][]int{}

	for len(queue) != 0 {
		tmp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp)
	}

	for i := 0; i < len(result)/2; i++ {
		tmp := result[i]
		result[i] = result[len(result)-i-1]
		result[len(result)-i-1] = tmp
	}

	return result
}
