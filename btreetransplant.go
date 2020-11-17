package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Parent != nil {
		rplc.Parent = node.Parent
		if rplc.Data < node.Parent.Data {
			node.Parent.Left = rplc
		} else {
			node.Parent.Right = rplc
		}
	} else {
		root = rplc
		rplc.Left = root.Left
		rplc.Right = root.Right
		return root
	}
	return root
}
