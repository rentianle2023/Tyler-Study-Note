package main

import (
	"fmt"
	"testing"
)

// 基准测试通常在"_test.go"文件中，并以 `Benchmark` 开头命名。
// `testing` 运行器多次执行每个基准测试函数，并在每次运行时增加 `b.N`，
// 直到它收集到精确的测量值。
func BenchmarkIntMin(b *testing.B) {
	// 通常，基准测试运行一个函数，我们在一个 `b.N` 次的循环内进行基准测试。
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

// 通常编写一个名称以 `Test` 开头的函数来创建测试。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` 会报告测试失败的信息，然后继续运行测试。
		// `t.Fatal*` 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 单元测试可以重复，所以会经常使用 *表驱动* 风格编写单元测试，
// 表中列出了输入数据，预期输出，使用循环，遍历并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run 可以运行一个 "subtests" 子测试，一个子测试对应表中一行数据。
		// 运行 `go test -v` 时，他们会分开显示。
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
