package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	input := []int{1, 1, 1, 1, 1, 0, 1}
	fmt.Println(isUnivalTree())
}

func isUnivalTree(root *TreeNode) bool {

}
