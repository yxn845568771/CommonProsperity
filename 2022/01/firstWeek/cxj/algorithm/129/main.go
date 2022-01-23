package main

import "fmt"

/*
title:求根节点到叶节点数字之和
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
eg1:
输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25
eg2:
输入：root = [4,9,0,5,1]
输出：1026
解释：
从根到叶子节点路径 4->9->5 代表数字 495
从根到叶子节点路径 4->9->1 代表数字 491
从根到叶子节点路径 4->0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026
*/
func main() {
	eg1 := &treeNode{
		Val: 1,
		Left: &treeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &treeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(sumNumbers(eg1)) // output : 25
	eg2 := &treeNode{
		Val: 4,
		Left: &treeNode{
			Val:   9,
			Left:  &treeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &treeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &treeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(sumNumbers(eg2)) // output : 1026
}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func sumNumbers(root *treeNode) int {
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(root *treeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	dfs(root.Left, sum, res)
	dfs(root.Right, sum, res)
}
