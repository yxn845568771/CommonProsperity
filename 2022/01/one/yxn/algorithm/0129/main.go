package _129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num int

func sumNumbers(root *TreeNode) int {
	num = 0
	if root == nil {
		return 0
	}
	d(root, 0)
	return num
}

func d(root *TreeNode, n int) {
	a := n*10 + root.Val
	if root.Left == nil && root.Right == nil {
		
		num += a
		return
	}
	if root.Left != nil {
		d(root.Left, a)
	}
	if root.Right != nil {
		d(root.Right, a)
	}
}
