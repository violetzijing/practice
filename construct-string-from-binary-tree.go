package main

func main() {
	node4 := &TreeNode{
		Val: 4,
	}
	node2 := &TreeNode{
		Val:  2,
		Left: node4,
	}
	node2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {

}
