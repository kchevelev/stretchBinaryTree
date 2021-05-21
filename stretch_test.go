package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestStretchK2(t *testing.T) {
	root := TreeNode{
		Val: 12,
		LeftNode: &TreeNode{
			Val:      81,
			LeftNode: nil,
			RightNode: &TreeNode{
				Val:       56,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &TreeNode{
			Val: 34,
			LeftNode: &TreeNode{
				Val:       19,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &TreeNode{
				Val:       6,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
	}

	k := 2

	var before string
	printTree(&root, "L", &before)
	stretch(&root, k)

	var after string
	expected := "L:6 L:6 L:40 L:40 R:28 R:28 R:17 R:17 L:9 L:9 R:3 R:3 "
	printTree(&root, "L", &after)
	if after != expected {
		t.Fatalf("expected %q,\n got %q", expected, after)
	}
}

func TestStretchK3(t *testing.T) {
	root := TreeNode{
		Val: 12,
		LeftNode: &TreeNode{
			Val:      81,
			LeftNode: nil,
			RightNode: &TreeNode{
				Val:       56,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &TreeNode{
			Val: 34,
			LeftNode: &TreeNode{
				Val:       19,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &TreeNode{
				Val:       6,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
	}

	k := 3

	var before string
	printTree(&root, "L", &before)
	stretch(&root, k)

	var after string
	expected := "L:4 L:4 L:4 L:27 L:27 L:27 R:18 R:18 R:18 R:11 R:11 R:11 L:6 L:6 L:6 R:2 R:2 R:2 "
	printTree(&root, "L", &after)
	if after != expected {
		t.Fatalf("expected %q,\n got %q", expected, after)
	}
}

func printTree(node *TreeNode, dir string, out *string) {

	*out += fmt.Sprintf("%v:%v ", dir, node.Val)

	if node.LeftNode != nil {
		printTree(node.LeftNode, "L", out)
	}

	if node.RightNode != nil {
		printTree(node.RightNode, "R", out)
	}
}
