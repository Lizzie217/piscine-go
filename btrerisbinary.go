package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, "", "", false, false)
}

func isBST(node *TreeNode, min, max string, hasMin, hasMax bool) bool {
	if node == nil {
		return true
	}

	if hasMin && node.Data <= min {
		return false
	}

	if hasMax && node.Data >= max {
		return false
	}

	return isBST(node.Left, min, node.Data, hasMin, true) &&
		isBST(node.Right, node.Data, max, true, hasMax)
}
