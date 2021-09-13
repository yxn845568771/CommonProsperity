package _226_invert_binary_tree

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
func invertTree(root *TreeNode) *TreeNode {
	if root==nil {
		return root
	}
	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)
	return root
}
