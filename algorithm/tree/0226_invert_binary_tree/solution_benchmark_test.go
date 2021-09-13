package _226_invert_binary_tree

import "testing"

var arg = tests[0].args

func Benchmark_isSameTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSameTree(arg.q, arg.p)
	}
}
