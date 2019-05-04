package main

func main() {
	node5 := &TreeNode{
		Val: 5,
	}
	node6 := &TreeNode{
		Val: 6,
	}
	node3 := &TreeNode{
		Val: 3,
		Kids: []*TreeNode{
			node5,
			node6,
		},
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node1 := &TreeNode{
		Val: 1,
		Kids: []*TreeNode{
			node3,
			node2,
			node4,
		},
	}
	fmt.Println(naryTreeLevelOrderTraversal(node1))
}

type TreeNode struct {
	Val  int
	Kids []*TreeNode
}

func naryTreeLevelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	result := [][]int{}

	node := root
	listSet := []*TreeNode{}
	for node != nil {
		result = append(result, node.Val)
		for len(node.Kids) != 0 {

		}
	}
}
