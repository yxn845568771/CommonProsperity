package _101_symmetric_tree

import (
	"test/algorithm/tree"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
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
			args: args{root: tree.NewBaseTreeNode(1, 2, 2, 3, 4, 4, 3)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
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
		args: args{root: tree.NewBaseTreeNode(1, 2, 2, 3, 4, 4, 3)},
		want: true,
	},
	{
		name: "1",
		args: args{root: tree.NewBaseTreeNode(1, 2, 2, 0, 3, 0, 3)},
	},
}
