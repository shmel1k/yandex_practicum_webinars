package scheduler

import "container/heap"

type task struct {
	takeTime int64
	taskID   string
}

type myContainerHeap struct {
	data []task
}

func (m *myContainerHeap) add(x int) {
	heap.Push(m, x)
}

func (m *myContainerHeap) pop() task {
	return heap.Pop(m).(task)
}

func (m *myContainerHeap) Push(x any) {
	m.data = append(m.data, x.(task))
}

func (m *myContainerHeap) Pop() any {
	x := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]
	return x
}

func (m *myContainerHeap) Top() task {
	return m.data[0]
}

func (m *myContainerHeap) Len() int {
	return len(m.data)
}

func (m *myContainerHeap) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *myContainerHeap) Less(i, j int) bool {
	return m.data[i].takeTime < m.data[j].takeTime
}
