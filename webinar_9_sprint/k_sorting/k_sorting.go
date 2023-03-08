package k_sorting

import (
	"container/heap"
	"sort"
)

/*
Есть массив из N элементов. Нужно найти K-тый элемент по убыванию(или возрастанию)
*/

func FindKthMaximumSort(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}

	sort.Ints(arr)
	return arr[len(arr)-k]
}

type myHeap struct {
	data []int
	idx  int
}

func (m *myHeap) Push(x any) {
	m.data[m.idx] = x.(int)
	m.idx++
}

func (m *myHeap) Pop() any {
	x := m.data[m.idx-1]
	m.idx--
	return x
}

func (m *myHeap) Less(i, j int) bool {
	return m.data[i] < m.data[j]
}

func (m *myHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *myHeap) Len() int {
	return m.idx
}

// Less
// Len
// Swap

func FindKthMaximumHeap(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}

	h := &myHeap{
		data: make([]int, k+1),
	}
	// [7, 9, 2, 5, 6, 3], k = 3
	//                 ^
	// [2, 9, 7, 5] -> 4???
	// Pop()
	// [5, 9, 7, 6]
	// Pop()
	// [6, 9, 7]
	// [3, 6, 9, 7]
	// Pop()
	// [6, 9, 7]
	// 6 -> result
	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Insert -> O(logN), где N -- размер кучи.
	// Конкретно в нашем случае -- O(logK)
	// O(NlogK) вместо O(NlogN)

	return heap.Pop(h).(int)
}

// На вход идут данные
// В любой момент времени мы хотим получать минимальный элемент

type myContainer struct {
	data []int
}

func (m *myContainer) add(x int) {
	m.data = append(m.data, x)
	sort.Ints(m.data)
}

func (m *myContainer) min() int {
	return m.data[0]
}

func (m *myContainer) pop() {
	m.data = m.data[1:]
}

type myContainerHeap struct {
	data []int
}

func (m *myContainerHeap) add(x int) {
	heap.Push(m, x)
}

func (m *myContainerHeap) min() int {
	return m.data[0]
}

func (m *myContainerHeap) pop() {
	heap.Pop(m)
}

func (m *myContainerHeap) Push(x any) {
	m.data = append(m.data, x.(int))
}

func (m *myContainerHeap) Pop() any {
	x := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return x
}

func (m *myContainerHeap) Len() int {
	return len(m.data)
}

func (m *myContainerHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *myContainerHeap) Less(i, j int) bool {
	return m.data[i] < m.data[j]
}
