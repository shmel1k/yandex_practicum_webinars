package benchmark

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sqrt5 = math.Sqrt(5)
)

func TestTwoEqualToTargetLine(t *testing.T) {
	testData := []struct {
		testName string
		nums     []int
		target   int
		expected bool
	}{
		{
			testName: "empty nums",
			nums:     []int{},
			target:   42,
			expected: false,
		},
		{
			testName: "42",
			nums:     []int{1, 10, 2, 15, 6, 27},
			target:   42,
			expected: true,
		},
	}

	for _, v := range testData {
		v := v
		t.Run(v.testName, func(t *testing.T) {
			got := TwoEqualToTargetLine(v.nums, v.target)
			assert.Equal(t, v.expected, got)
		})
	}
}

func TestTwoEqualToTargetSquared(t *testing.T) {
	testData := []struct {
		testName string
		nums     []int
		target   int
		expected bool
	}{
		{
			testName: "empty nums",
			nums:     []int{},
			target:   42,
			expected: false,
		},
		{
			testName: "42",
			nums:     []int{1, 10, 2, 15, 6, 27},
			target:   42,
			expected: true,
		},
	}

	for _, v := range testData {
		v := v
		t.Run(v.testName, func(t *testing.T) {
			got := TwoEqualToTargetSquared(v.nums, v.target)
			assert.Equal(t, v.expected, got)
		})
	}
}

func prepareTwoEqualToTarget(count int) []int {
	res := make([]int, 0, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(1000000)
		res = append(res, n)
	}
	return res
}

func BenchmarkTwoEqualToTarget(b *testing.B) {
	data := prepareTwoEqualToTarget(1000000)
	b.ResetTimer()

	b.Run("squared", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = TwoEqualToTargetSquared(data, 8423)
		}
	})
	b.Run("line", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = TwoEqualToTargetLine(data, 8423)
		}
	})
}

func FiboRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FiboRecursive(n-2) + FiboRecursive(n-1)
}

func FiboOptimized(n int) int {
	if n <= 1 {
		return n
	}
	fibo := make([]int, n+1)
	fibo[1] = 1
	for i := 2; i <= n; i++ {
		fibo[i] = fibo[i-2] + fibo[i-1]
	}
	return fibo[n]
}

func FiboUltraOptimized(n int) int {
	if n <= 1 {
		return n
	}
	first := 0
	second := 1
	tmp := 0
	for i := 2; i <= n; i++ {
		tmp = second + first
		first = second
		second = tmp
	}
	return second
}

func FiboProbablyOptimized(n int) float64 {
	return math.Pow((1+sqrt5)/2, 30) - math.Pow((1-sqrt5)/2, 30)/sqrt5
}

func TestFiboUltraOptimized(t *testing.T) {
	testData := []struct {
		testName string
		nth      int
		expected int
	}{
		{
			testName: "first",
			nth:      1,
			expected: 1,
		},
		{
			testName: "second",
			nth:      2,
			expected: 1,
		},
		{
			testName: "third",
			nth:      3,
			expected: 2,
		},
		{
			testName: "tenth",
			nth:      10,
			expected: 55,
		},
	}

	for _, v := range testData {
		v := v
		t.Run(v.testName, func(t *testing.T) {
			got := FiboUltraOptimized(v.nth)
			assert.Equal(t, v.expected, got)
		})
	}
}

func BenchmarkFibo(b *testing.B) {
	count := 1000
	// NOTE(shmel1k@): uncomment if you want to test Recursive.
	//	b.Run("recursive", func(b *testing.B) {
	//		for i := 0; i < b.N; i++ {
	//			FiboRecursive(count)
	//		}
	//	})
	b.Run("optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FiboOptimized(count)
		}
	})
	b.Run("ultra_optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FiboUltraOptimized(count)
		}
	})
	b.Run("probably_optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FiboProbablyOptimized(count)
		}
	})
}
