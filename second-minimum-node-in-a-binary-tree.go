package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	stack := []*TreeNode{}
	node := root
	min := 2147483647
	secMin := 2147483647
	max := root.Val
	for node != nil || len(stack) != 0 {
		for node != nil {
			if min > node.Val {
				secMin = min
				min = node.Val
			} else if node.Val != min && node.Val < secMin {
				secMin = node.Val
			}
			if max < node.Val {
				max = node.Val
			}
			stack = append(stack, node)
			node = node.Left
		}
		if node == nil && len(stack) != 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}

	if min == secMin {
		return -1
	}
	if max != secMin && secMin == 2147483647 {
		return -1
	}

	return secMin
}
