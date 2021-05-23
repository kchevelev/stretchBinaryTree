package data

func GetSampleTree() TreeNode {
	return TreeNode{
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
}
