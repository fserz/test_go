package testRun

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("test main start executing...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")
	code := m.Run()
	fmt.Println("test main finished")
	os.Exit(code)
}

// 普通参数化测试
func TestMultiply(t *testing.T) {
	type testClass struct {
		name     string
		x, y     int
		expected int
	}
	cases := []testClass{
		{"both positive", 2, 3, 6},
		{"both negative", -2, -3, 6},
		{"one zero", 0, 10, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.x * c.y; got != c.expected {
				t.Errorf("expected %d, got %d", c.expected, got)
			}
		})
	}
}

// 对于复杂逻辑，测试可能会有多个层级，t.Run 支持嵌套使用。
func TestMathOperations(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		if 1+2 != 3 {
			t.Error("1 + 2 should equal 3")
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		t.Run("Positive", func(t *testing.T) {
			if 2*3 != 6 {
				t.Error("2 * 3 should equal 6")
			}
		})

		t.Run("Negative", func(t *testing.T) {
			if (-2)*3 != -6 {
				t.Error("-2 * 3 should equal -6")
			}
		})
	})

	t.Run("Error test", func(t *testing.T) {
		if 2*3 != 4 {
			t.Error("2 * 3 != 5 error~")
		}
	})
}

func TestParallel(t *testing.T) {
	tests := []struct {
		name  string
		delay int // 模拟延迟
	}{
		{"Test1", 1},
		{"Test2", 2},
		{"Test3", 3},
	}

	for _, tt := range tests {
		tt := tt // 必须重新声明避免闭包问题
		t.Run(tt.name, func(t *testing.T) {
			// 开启并行，一旦标记为并行，该子测试将在单独的 Goroutine 中运行，与其他标记为并行的子测试同时执行。
			t.Parallel()
			// 模拟逻辑，例如：time.Sleep(time.Duration(tt.delay) * time.Second)
			time.Sleep(time.Duration(tt.delay) * time.Second)
			t.Logf("Running %s with delay %d", tt.name, tt.delay)
		})
	}
}

func TestParallel1(t *testing.T) {
	tests := []struct {
		name  string
		delay int
	}{
		{"Test1", 1},
		{"Test2", 2},
		{"Test3", 3},
		{"Test4", 1},
		{"Test5", 2},
		{"Test6", 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			time.Sleep(time.Duration(tt.delay) * time.Second)
			t.Logf("Running %s with delay %d", tt.name, tt.delay)
		})
	}
}

// 测试是否在一个goroutine
func TestParallel2(t *testing.T) {
	tests := []struct {
		name  string
		delay int
	}{
		{"Test1", 1},
		{"Test2", 2},
		{"Test3", 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("Running %s in Goroutine %d", tt.name, getGoroutineID())
		})
	}
}

// 获取当前 Goroutine ID（调试用）
func getGoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idField)
	return id
}

func BenchmarkS(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
