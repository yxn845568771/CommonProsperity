package _110_balanced_binary_tree

import (
	"test/algorithm/tree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{root: tree.NewBaseTreeNode(3, 9, 20, 0, 0, 15, 7)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{root: tree.NewBaseTreeNode(3, 9, 20, 0, 0, 15, 7)},
		want: true,
	},
	{
		name: "2",
		args: args{root: tree.NewBaseTreeNode(1,2,2,3,3,0,0,4,4)},
		want: false,
	},
	{
		name: "3",
		args: args{root: tree.NewBaseTreeNode()},
		want: true,
	},
}
