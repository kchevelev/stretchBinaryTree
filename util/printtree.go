package util

import (
	"fmt"

	"github.com/kchevelev/stretchbinarytree/data"
)

func FlattenTree(node *data.TreeNode, dir string, out *string) {

	*out += fmt.Sprintf("%v:%v ", dir, node.Val)

	if node.LeftNode != nil {
		FlattenTree(node.LeftNode, "L", out)
	}

	if node.RightNode != nil {
		FlattenTree(node.RightNode, "R", out)
	}
}
