package _101_symmetric_tree

import "testing"

var arg = tests[0].args

func Benchmark_isSymmetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isSymmetric(arg.root)
	}
}
