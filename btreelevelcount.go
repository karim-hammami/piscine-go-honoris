package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return TreeIterativeCounter(root, 1)
}

func TreeIterativeCounter(root *TreeNode, n int) int {
	if root == nil {
		return n - 1
	}

	if TreeIterativeCounter(root.Left, n+1) > TreeIterativeCounter(root.Right, n+1) {
		return TreeIterativeCounter(root.Left, n+1)
	}
	return TreeIterativeCounter(root.Right, n+1)
}
