package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp := map[int]bool{}
			tmp[node.Val] = true
			if node.Left != nil && node.Right != nil {
				if (node.Left.Val == x && node.Right.Val == y) ||
					(node.Left.Val == y && node.Right.Val == x) {
					return false
				}
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if tmp[x] && tmp[y] {
			return true
		}
	}

	return false
}
