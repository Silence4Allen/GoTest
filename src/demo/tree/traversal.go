package tree

import "fmt"

func (node *Node) Traversal() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
