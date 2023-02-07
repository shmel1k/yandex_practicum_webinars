package main

import (
	"fmt"
	"sync"
	"time"
)

func wrongMutexUsage(mu sync.Mutex, mp map[string]string) {
	mu.Lock()
	mp["abacaba"] = fmt.Sprintf("%s", time.Now())
	mu.Unlock()
}

func rightMutexUsage(mu sync.Mutex, mp map[string]string) {
	now := fmt.Sprintf("%s", time.Now())
	mu.Lock()
	mp["abacaba"] = now
	mu.Unlock()
}

func wrongPreallocation(size int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
	return arr
}

func rightPreallocation(size int) []int {
	arr := make([]int, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
	return arr
}

func wrongStringAppend() string {
	a := "abacaba"
	b := "abacaba"
	for i := 0; i < 1000; i++ {
		a += b
	}
	return a
}

func rightStringAppend() string {
	a := []byte("abacaba")
	b := "abacaba"
	for i := 0; i < 1000; i++ {
		a = append(a, b...)
	}
	return string(a)
}

func main() {
	_ = wrongMutexUsage
	_ = rightMutexUsage
	_ = wrongStringAppend
	_ = rightStringAppend
}
