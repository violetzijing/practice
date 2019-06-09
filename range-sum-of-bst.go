package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if node.Val >= L && node.Val <= R {
				sum += node.Val
			}
			if node.Val > L {
				queue = append(queue, node.Left)
			}
			if node.Val < R {
				queue = append(queue, node.Right)
			}
		}
	}

	return sum
}
