package _100_same_tree

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
				p: tree.NewBaseTreeNode(1, 0, 2),
				q: tree.NewBaseTreeNode(1, 0, 2),
			},
			want: true,
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
			p: tree.NewBaseTreeNode(1,0,2),
			q: tree.NewBaseTreeNode(1,0,2),
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			p: tree.NewBaseTreeNode(1,0,2),
			q: tree.NewBaseTreeNode(1,0,1),
		},
		want: false,
	},
	{
		name: "2",
		args: args{
			p: tree.NewBaseTreeNode(1,0,1,0,1),
			q: tree.NewBaseTreeNode(1,0,1,1),
		},
		want: false,
	},
}
