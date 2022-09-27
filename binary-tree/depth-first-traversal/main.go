package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func preorder(node *Node) {
	if node == nil {
		return
	}

	// 1. Traverse root
	fmt.Printf("%v ", node.data)

	// 2. Pre-order on left sub-tree.
	preorder(node.left)

	// 3. Pre-order on right sub-tree.
	preorder(node.right)
}

func inorder(node *Node) {
	if node == nil {
		return
	}

	// 1. In-order on left sub-tree.
	inorder(node.left)

	// 2. Traverse root.
	fmt.Printf("%v ", node.data)

	// 3. In-order on right sub-tree.
	inorder(node.right)
}

func postorder(node *Node) {
	if node == nil {
		return
	}
	// 1. Post-order on left sub-tree.
	postorder(node.left)

	// 2. Post-order on right sub-tree.
	postorder(node.right)

	// 3. Traverse root.
	fmt.Printf("%v ", node.data)
}

func main() {
	tree := NewNode(0)
	tree.left = NewNode(1)
	tree.right = NewNode(2)
	tree.left.left = NewNode(3)
	tree.left.right = NewNode(4)

	// print binary tree
	fmt.Println("Given Binary Tree")
	fmt.Println("   0  ")
	fmt.Println("  / \\ ")
	fmt.Println(" 1   2")
	fmt.Println("/ \\   ")
	fmt.Println("3 4   ")

	// pre order traversal
	fmt.Println("\nPre Order Traversal")
	preorder(tree)
	fmt.Println()

	// in order traversal
	fmt.Println("\nIn Order Traversal")
	inorder(tree)
	fmt.Println()

	// post order traversal
	fmt.Println("\nPost Order Traversal")
	postorder(tree)
	fmt.Println()
}
