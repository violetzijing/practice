package main

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func main() {

}

func max(tree *Node) int {
	if tree.Right != nil {
		return tree.Value
	}

	return max(tree.Right)
}
