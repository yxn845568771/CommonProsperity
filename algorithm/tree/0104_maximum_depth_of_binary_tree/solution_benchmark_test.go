package _104_maximum_depth_of_binary_tree

import "testing"

var arg = tests[0].args

func Benchmark_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxDepth(arg.root)
	}
}
