package _094_binary_tree_inorder_traversal

import "test/algorithm/tree"

type TreeNode = tree.BaseTreeNode

func inorderTraversal(root *TreeNode) []int {
	var a []int
	inOrder(root, &a)
	return a
}

func inOrder(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, a)
	*a = append(*a, root.Val)
	inOrder(root.Right, a)
}
