package _100_same_tree

import "test/algorithm/tree"

type TreeNode = tree.BaseTreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left,q.Left){
		return false
	}
	return isSameTree(p.Right,q.Right)
}
