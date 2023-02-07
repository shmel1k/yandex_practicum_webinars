package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.HandleFunc("/cpu_throttle", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		for {
			if time.Since(start) >= 10*time.Second {
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
