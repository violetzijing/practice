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
	fmt.Println(averageOfLevels(node3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	result := []float64{}
	queue := []*TreeNode{
		root,
	}

	for len(queue) != 0 {
		sum := float64(0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			sum += float64(node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, sum/float64(size))
	}

	return result
}
