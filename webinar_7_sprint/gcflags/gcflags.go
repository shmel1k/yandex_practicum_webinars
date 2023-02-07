package main

import (
	"fmt"
	"time"
)

func MyVeryOwnFunc(arr []int) {
	for i := 0; i < len(arr); i++ {
		v := len(arr)
		if v > 3 {
			fmt.Println(v)
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	go MyVeryOwnFunc(arr)
	time.Sleep(time.Second)
}
