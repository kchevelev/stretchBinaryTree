package main

import (
	"fmt"

	"github.com/kchevelev/stretchbinarytree/data"
	"github.com/kchevelev/stretchbinarytree/util"
)

func main() {
	var root data.TreeNode
	var k int
	var before, after string

	/* case 1 */

	// stretch the tree by a factor of 2
	root = data.GetSampleTree()
	k = 2

	// print out the state before
	before = ""
	util.FlattenTree(&root, "L", &before)
	fmt.Printf("\nNodes before the stretch: %q\n", before)

	// perform the stretch
	fmt.Printf("Stretching by a factor of %v ...\n", k)
	stretch(&root, k)

	// print out the state after
	after = ""
	util.FlattenTree(&root, "L", &after)
	fmt.Printf("Nodes after the stretch: %q\n", after)

	/* case 2 */

	// stretch the tree by a factor of 3
	root = data.GetSampleTree()
	k = 3

	// print out the state before
	before = ""
	util.FlattenTree(&root, "L", &before)
	fmt.Printf("\nNodes before the stretch: %q\n", before)

	// perform the stretch
	fmt.Printf("Stretching by a factor of %v ...\n", k)
	stretch(&root, k)

	// print out the state after
	after = ""
	util.FlattenTree(&root, "L", &after)
	fmt.Printf("Nodes after the stretch: %q\n", after)

}
