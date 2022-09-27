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

func levelorder(node *Node) {
	// 1. Initialize an empty queue.
	queue := []*Node{}

	// 2. Set root as temp.
	queue = append(queue, node)

	// 3. Loop.
	for len(queue) > 0 {
		// set the first item as tempNode
		tempNode := queue[0]

		// remove the first item from queue
		queue = queue[1:]

		// print tempNode value
		fmt.Printf("%v ", tempNode.data)

		// add left child to the queue
		if tempNode.left != nil {
			queue = append(queue, tempNode.left)
		}

		// add right child to the queue
		if tempNode.right != nil {
			queue = append(queue, tempNode.right)
		}
	}
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

	// level order traversal
	fmt.Println("\nLevel Order Traversal")
	levelorder(tree)
	fmt.Println()
}
