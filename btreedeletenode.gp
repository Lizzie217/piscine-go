package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Case 1: no left child
	if node.Left == nil {
		return BTreeTransplant(root, node, node.Right)
	}

	// Case 2: no right child
	if node.Right == nil {
		return BTreeTransplant(root, node, node.Left)
	}

	// Case 3: two children
	successor := BTreeMin(node.Right)

	// If successor is not direct right child
	if successor.Parent != node {
		// replace successor with its right child
		root = BTreeTransplant(root, successor, successor.Right)

		// move node.Right under successor
		successor.Right = node.Right
		if successor.Right != nil {
			successor.Right.Parent = successor
		}
	}

	// replace node with successor
	root = BTreeTransplant(root, node, successor)

	// attach left subtree
	successor.Left = node.Left
	if successor.Left != nil {
		successor.Left.Parent = successor
	}

	return root
}
