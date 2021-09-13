package _110_balanced_binary_tree

import "testing"

var arg = tests[0].args

func Benchmark_isBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isBalanced(arg.root)
	}
}

