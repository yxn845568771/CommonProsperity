package _226_invert_binary_tree

import (
	"test/algorithm/tree"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				p: tree.NewBaseTreeNode(4, 7, 2, 9, 6, 3, 1),
				q: invertTree(tree.NewBaseTreeNode(4, 2, 7, 1, 3, 6, 9)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	p *TreeNode
	q *TreeNode
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			p: tree.NewBaseTreeNode(4, 7, 2, 9, 6, 3, 1),
			q: invertTree(tree.NewBaseTreeNode(4, 2, 7, 1, 3, 6, 9)),
		},
		want: false,
	},
	{
		name: "2",
		args: args{
			p: tree.NewBaseTreeNode(),
			q: invertTree(tree.NewBaseTreeNode()),
		},
		want: true,
	},
	{
		name: "3",
		args: args{
			p: tree.NewBaseTreeNode(2,3,1),
			q: invertTree(tree.NewBaseTreeNode(2,1,3)),
		},
		want: true,
	},
}
