package k_sorting

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	testName string
	arr      []int
	k        int
	expected int
}{
	{
		testName: "empty array",
		arr:      []int{},
		k:        42,
		expected: -1,
	},
	{
		testName: "first element",
		arr: []int{
			5, 4, 3, 2, 1,
		},
		k:        1,
		expected: 5,
	},
	{
		testName: "last element",
		arr: []int{
			5, 4, 3, 1, 2, 9, 8, 7,
		},
		k:        1,
		expected: 9,
	},
}

func TestFindKthMaximumSort(t *testing.T) {
	for _, v := range testData {
		v := v
		t.Run(v.testName, func(t *testing.T) {
			got := FindKthMaximumSort(v.arr, v.k)
			assert.Equal(t, v.expected, got)
		})
	}
}

func TestFindKthMaximumHeap(t *testing.T) {
	for _, v := range testData {
		v := v
		t.Run(v.testName, func(t *testing.T) {
			got := FindKthMaximumHeap(v.arr, v.k)
			assert.Equal(t, v.expected, got)
		})
	}
}

func generateBenchmarkData(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, rand.Intn(10000))
	}
	return result
}

func BenchmarkTestFindKthMaximumSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := generateBenchmarkData(10000)
		b.StartTimer()
		FindKthMaximumSort(data, 1)
	}
}

func BenchmarkTestFindKthMaximumHeap(b *testing.B) {
	b.StopTimer()
	data := generateBenchmarkData(10000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data1 := make([]int, 10000)
		copy(data1, data)
		b.StartTimer()
		FindKthMaximumHeap(data1, 100)
	}
}

func BenchmarkMyContainerNoHeap(b *testing.B) {
	data := generateBenchmarkData(40000)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := &myContainer{
			data: make([]int, 0, len(data)),
		}
		b.StartTimer()

		for j := 0; j < len(data); j++ {
			if (j+1)%10 == 0 {
				m.pop()
			}
			m.add(data[j])
		}
	}
}

func BenchmarkMyContainerHeap(b *testing.B) {
	data := generateBenchmarkData(40000)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := &myContainerHeap{
			data: make([]int, 0, len(data)),
		}
		b.StartTimer()

		for j := 0; j < len(data); j++ {
			if (j+1)%10 == 0 {
				m.pop()
			}
			m.add(data[j])
		}
	}
}
