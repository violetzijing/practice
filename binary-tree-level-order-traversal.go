package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	size := len(queue)
	result := [][]int{}

	for size != 0 {
		levelResult := []int{}
		size = len(queue)
		for i := 0; i < size; i++ {
			levelResult = append(levelResult, queue[i].Val)
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		if len(levelResult) != 0 {
			result = append(result, levelResult)
		}

	}

	return result
}
