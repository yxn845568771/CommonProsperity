package _095_unique_binary_search_trees_2

import (
	"test/algorithm/tree"
)

type TreeNode = tree.BaseTreeNode

var length int

var root *TreeNode

var trees []*TreeNode

func copyTree(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	tr := &TreeNode{
		Val: t.Val,
	}
	tr.Left = copyTree(t.Left)
	tr.Right = copyTree(t.Right)
	return tr
}

func generateTrees(n int) []*TreeNode {
	length = n
	if n <= 0 {
		return nil
	}
	tr := &TreeNode{
		Val: 1,
	}
	root = tr
	addNode(1, tr)
	return trees
}

func addNode(n int, tree *TreeNode) {
	tree.Val = n
	if n == length {
		trees = append(trees, copyTree(root))
		tree = nil
		return
	}
	tree.Left = &TreeNode{}
	tree.Right = &TreeNode{}
	addNode(n+1, tree.Left)
	addNode(n+1, tree.Right)
}
