package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	stack := []*TreeNode{}
	result1 := []int{}
	result2 := []int{}

	node := root1
	for node != nil || len(stack) != 0 {
		for node != nil {
			if node.Left == nil && node.Right == nil {
				result1 = append(result1, node.Val)
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

	node = root2
	for node != nil || len(stack) != 0 {
		for node != nil {
			if node.Left == nil && node.Right == nil {
				result2 = append(result2, node.Val)
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
	fmt.Println(result1)
	fmt.Println(result2)
	if len(result1) != len(result2) {
		return false
	}

	for i := 0; i < len(result1); i++ {
		if result1[i] != result2[i] {
			return false
		}
	}

	return true
}
