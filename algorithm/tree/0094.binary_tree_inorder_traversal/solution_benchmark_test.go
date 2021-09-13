package _094_binary_tree_inorder_traversal

import (
	"testing"
)

var arg = tests[0].args

func Benchmark_inorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		inorderTraversal(arg.root)
	}
}