package main

import (
	"log"
)

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func stretch(root *TreeNode, stretchAmount int) {
	if stretchAmount <= 0 {
		log.Fatalf("invalid stretchAmount value")
	}
	nodeType := "left"
	if root != nil {
		keepStretching(root, stretchAmount, nodeType)
	}
}

func keepStretching(node *TreeNode, stretchAmount int, nodeType string) {
	// fmt.Printf("\ngot: %+v\n", node)

	// keep references for after injection
	leftRef := node.LeftNode
	rightRef := node.RightNode

	// clear existing references on node being stretched
	node.LeftNode = nil
	node.RightNode = nil

	// calculate stretched value
	val := node.Val / stretchAmount

	// fmt.Printf("updating current node value to %v\n", val)
	// update value of node being stretched
	node.Val = val

	// stretch a node N times
	for i := 1; i < stretchAmount; i++ {
		switch nodeType {
		case "left":
			// fmt.Println("stretching left")
			node.LeftNode = &TreeNode{
				Val:       val,
				LeftNode:  nil,
				RightNode: nil,
			}
			// fmt.Printf("new left node: %+v\n", node.LeftNode)
			node = node.LeftNode
			// fmt.Printf("advancing to left node: %+v\n", node)

		case "right":
			// fmt.Println("stretching right")
			node.RightNode = &TreeNode{
				Val:       val,
				LeftNode:  nil,
				RightNode: nil,
			}
			// fmt.Printf("new right node: %+v\n", node.RightNode)
			node = node.RightNode
			// fmt.Printf("advancing to right node: %+v\n", node)
		}
	}

	// restore references
	node.LeftNode = leftRef
	node.RightNode = rightRef
	// fmt.Printf("restoring original references: %+v\n", node)

	if node.LeftNode != nil {
		// fmt.Printf("processing left node: %+v\n", node.LeftNode)
		keepStretching(node.LeftNode, stretchAmount, "left")
	}

	if node.RightNode != nil {
		// fmt.Printf("processing right node: %+v\n", node.RightNode)
		keepStretching(node.RightNode, stretchAmount, "right")
	}
}
