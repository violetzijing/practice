package main

// https://leetcode.com/problems/merge-two-binary-trees/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node15 := &TreeNode{
		Val: 5,
	}
	node13 := &TreeNode{
		Val:  3,
		Left: node15,
	}
	node12 := &TreeNode{
		Val: 2,
	}
	node11 := &TreeNode{
		Val:   1,
		Left:  node13,
		Right: node12,
	}
	tree1 := node11

	node24 := &TypeNode{
		Val: 4,
	}
	node27 := &TypeNode{
		Val: 7,
	}
	node21 := &TypeNode{
		Val:   1,
		Right: node24,
	}
	node23 := &TypeNode{
		Val:   3,
		Right: node27,
	}
	node22 := &TypeNode{
		Val:   2,
		Left:  node11,
		Right: node23,
	}

	tree2 := node22
	fmt.Println(mergeTrees(tree1, tree2))
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var node = t1
	stack := []*TreeNode{}

	for t1 != nil || len(stack) != 0 {
		node

	}

}
