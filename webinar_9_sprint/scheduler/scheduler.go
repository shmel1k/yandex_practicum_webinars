package scheduler

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
API:
UpdateHost(hostname) -> task_id
GetHostStatus(hostname) -> bool (true -- finished, false -- not finished)

for i := 0; i < 20000; i++ {
    UpdateHost(fmt.Sprintf("host%d"))
}

0. Мы создали Н задач в независимом(внешнем) сервисе. И нам нужно дождаться за-
вершения всех задач

Наш код:
1. На вход подаётся список задач. У каждой задачи есть ID.
2.
*/

type myChecker struct {
	tasks         *myContainerHeap
	queue         chan string
	mu            sync.RWMutex
	requestsCount uint64
	start         time.Time
}

func MyChecker(tasks []string) *myChecker {
	data := make([]task, 0, len(tasks))
	now := time.Now()
	for _, v := range tasks {
		data = append(data, task{
			takeTime: now.Unix() + rand.Int63n(2),
			taskID:   v,
		})
	}
	m := &myChecker{
		tasks: &myContainerHeap{
			data: data,
		},
		queue: make(chan string),
		start: time.Now(),
	}
	heap.Init(m.tasks)

	return m
}

func (m *myChecker) Wait() {
	m.runListeners(10)
	for {
		now := time.Now().Unix()
		tasks := make([]task, 0, 4)
		m.mu.Lock()
		for m.tasks.Len() > 0 && m.tasks.Top().takeTime < now {
			tasks = append(tasks, m.tasks.Pop().(task))
		}
		m.mu.Unlock()
		for _, t := range tasks {
			t := t
			go func() {
				m.addTaskToQueue(t.taskID)
			}()
		}
		if m.empty() {
			break
		}
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Waiting for tasks to finish, tasks left = ", len(m.tasks.data))
	}
	fmt.Printf("Got %d requests withing %s duration", m.requestsCount, time.Since(m.start))
}

func (m *myChecker) empty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.tasks.data) == 0
}

func (m *myChecker) addTaskToQueue(taskID string) {
	select {
	case m.queue <- taskID:
	}
}

func (m *myChecker) addTaskBackToHeap(t task) {
	m.mu.Lock()
	m.tasks.Push(t)
	m.mu.Unlock()
}

func (m *myChecker) runListeners(n int) {
	for i := 0; i < n; i++ {
		go func() {
			for {
				select {
				case t := <-m.queue:
					status := m.checkReadiness(t)
					if !status {
						now := time.Now()
						m.addTaskBackToHeap(task{
							takeTime: now.Unix() + rand.Int63n(2),
							taskID:   t,
						})
					}
				}
			}
		}()
	}
}

func (m *myChecker) checkReadiness(taskID string) bool {
	n := rand.Intn(5)
	atomic.AddUint64(&m.requestsCount, 1)
	return n%5 == 0
}
