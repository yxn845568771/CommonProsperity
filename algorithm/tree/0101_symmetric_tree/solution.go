package _101_symmetric_tree

import (
	"test/algorithm/tree"
)

type TreeNode = tree.BaseTreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return eq(root.Left,root.Right)
}

func eq(p,q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || q.Val != p.Val {
		return false
	}
	return eq(p.Left,q.Right) && eq(p.Right,q.Left)
}
