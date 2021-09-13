package _104_maximum_depth_of_binary_tree

import (
	"test/algorithm/tree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{root: tree.NewBaseTreeNode(3, 9, 20, 0, 0, 15, 17)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
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
	want int
}{
	{
		name: "1",
		args: args{root: tree.NewBaseTreeNode(3, 9, 20, 0, 0, 15, 17)},
		want: 3,
	},
	{
		name: "2",
		args: args{root: tree.NewBaseTreeNode(1, 0, 2)},
		want: 2,
	},
	{
		name: "3",
		args: args{root: tree.NewBaseTreeNode()},
		want: 0,
	},
	{
		name: "4",
		args: args{root: tree.NewBaseTreeNode(0)},
		want: 1,
	},
}
