package _094_binary_tree_inorder_traversal

import (
	"reflect"
	"test/algorithm/tree"
	"testing"
)

type args struct {
	root *TreeNode
}

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{root: tree.NewBaseTreeNode(1, 0, 2, 3)},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "testcase 1",
		args: args{
			root: tree.NewBaseTreeNode(1, 0, 2, 3),
		},
		want: []int{1, 3, 2},
	},
	{
		name: "testcase 2",
		args: args{
			root: tree.NewBaseTreeNode(),
		},
		want: nil,
	},
	{
		name: "testcase 3",
		args: args{
			root: tree.NewBaseTreeNode(1),
		},
		want: []int{1},
	},
	{
		name: "testcase 4",
		args: args{
			root: tree.NewBaseTreeNode(1, 2),
		},
		want: []int{2, 1},
	},
	{
		name: "testcase 5",
		args: args{
			root: tree.NewBaseTreeNode(1, 0, 2),
		},
		want: []int{1, 2},
	},
}
