package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	node := root
	i := 0
	result := 0
	i = i<<1 | root.Val
	if root.Left == nil && root.Left {
		result += i
	}

}
