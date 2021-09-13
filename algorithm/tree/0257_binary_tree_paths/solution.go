package _257_binary_tree_paths

import (
	"fmt"
	"test/algorithm/tree"
)

type TreeNode = tree.BaseTreeNode

func binaryTreePaths(root *TreeNode) []string {
	a := &as{S: []string{}}
	a.digui(root,"")
	return a.S
}

type as struct {
	S []string
}

func (a *as) digui(root *TreeNode, str string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		a.S = append(a.S, fmt.Sprintf("%s%d", str, root.Val))
	}
	a.digui(root.Left, fmt.Sprintf("%s%d->", str, root.Val))
	a.digui(root.Right, fmt.Sprintf("%s%d->", str, root.Val))
}
