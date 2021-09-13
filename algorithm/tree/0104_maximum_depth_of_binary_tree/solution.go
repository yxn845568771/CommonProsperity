package _104_maximum_depth_of_binary_tree

import "test/algorithm/tree"

type TreeNode = tree.BaseTreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getBigger(maxDepth(root.Left),maxDepth(root.Right))
}

func getBigger(a,b int) int{
	if a > b {
		return a
	}
	return b
}