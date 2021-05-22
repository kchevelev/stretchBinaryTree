package data

type NodeType string

const (
	LeftNode  NodeType = "left"
	RightNode NodeType = "right"
)

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}
