package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	if root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)

}
