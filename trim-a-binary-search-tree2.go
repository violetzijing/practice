package main

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	for root.Val < L || root.Val > R {
		if root.Val < L {
			root = root.Right
		}
		if root.Val > R {
			root = root.Left
		}
	}

	node := root
	for node != nil {
		for node.Left != nil && node.Left.Val < L {
			node.Left = node.Left.Right
		}
		node = node.Left
	}

	node = root
	for node != nil {
		for node.Right != nil && node.Right.Val > R {
			node.Right = node.Right.Left
		}
		node = node.Right
	}

	return root
}
