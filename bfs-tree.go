package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	fmt.Println(bfsTree(node3))
}

func bfsTree(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		result = append(result, node.Val)
		queue = queue[1:]
		if node.Left != nil {
			fmt.Println("left node: ", node.Left)
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			fmt.Println("right node: ", node.Right)
			queue = append(queue, node.Right)
			fmt.Println("last in queue: ", queue[len(queue)-1])
			fmt.Println("queue: ", queue[0])
		}
	}

	return result
}
