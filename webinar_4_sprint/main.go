package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

var password = "94535667"
var md []byte

func main() {
	h := md5.New()
	md = h.Sum([]byte(password))
	fmt.Printf("%x\n", md)
	t1 := time.Now()
	ch := make(chan int64)
	for i := int64(0); i < 100000000; i++ {
		go toMD5g(i, ch)
	}
	i := <-ch
	fmt.Println("Bingo, password ", i)
	fmt.Println(time.Since(t1))

}

func toMD5g(i int64, ch chan int64) {
	h := md5.New()
	var buf []byte
	buf = strconv.AppendInt(buf, int64(i), 16)
	if bytes.Equal(md, h.Sum(buf)) {
		ch <- i
		close(ch)
	}
}
