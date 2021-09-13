package _110_balanced_binary_tree

import "test/algorithm/tree"

type TreeNode = tree.BaseTreeNode

func isBalanced(root *TreeNode) bool {
	if depth(root) < 0 {
		return false
	}
	return true
}

func depth(root * TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	return getResult(left,right)
}

func getResult (a,b int) int {
	if a < 0 || b < 0 ||a - b > 1 || a - b < -1 {
		return  - 1
	}
	if a > b {
		return a +1
	}
	return b +1
}
