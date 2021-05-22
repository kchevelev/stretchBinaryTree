package main

import (
	"log"

	"github.com/kchevelev/stretchbinarytree/data"
)

func stretch(root *data.TreeNode, stretchAmount int) {
	if stretchAmount <= 0 {
		log.Fatalf("invalid stretchAmount value")
	}
	nodeType := data.LeftNode
	if root != nil {
		keepStretching(root, stretchAmount, nodeType)
	}
}

func keepStretching(node *data.TreeNode, stretchAmount int, nodeType data.NodeType) {
	// keep references for after injection
	leftRef := node.LeftNode
	rightRef := node.RightNode

	// calculate stretched value
	val := node.Val / stretchAmount

	// update value of node being stretched
	node.Val = val
	// clear existing references
	node.LeftNode = nil
	node.RightNode = nil

	// stretch a node N times
	for i := 1; i < stretchAmount; i++ {
		switch nodeType {
		case data.LeftNode:
			// stretching to the left
			node.LeftNode = &data.TreeNode{
				Val:       val,
				LeftNode:  nil,
				RightNode: nil,
			}
			// advancing to left node
			node = node.LeftNode

		case data.RightNode:
			// stretching to the right
			node.RightNode = &data.TreeNode{
				Val:       val,
				LeftNode:  nil,
				RightNode: nil,
			}
			// advancing to right node
			node = node.RightNode
		}
	}

	// restore references on the last instance of stretched node
	node.LeftNode = leftRef
	node.RightNode = rightRef

	if node.LeftNode != nil {
		// recursively keep processing node to the left
		keepStretching(node.LeftNode, stretchAmount, data.LeftNode)
	}

	if node.RightNode != nil {
		// recursively keep processing node to the right
		keepStretching(node.RightNode, stretchAmount, data.RightNode)
	}
}
