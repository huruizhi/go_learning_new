package split

import (
	"testing"
)

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a,b,c,d", ",")
	}
}

// 基准测试 go test -bench=Split
// 基准测试内存 go test -bench=Split --benchmem
// 基准测试 go test -bench=. 执行所有的基准测试
