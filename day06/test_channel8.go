package main

import (
	"fmt"
	"time"
)

// 用 select 实现超时 (timeout)。

func main() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout.")
		}
		w <- true
	}()
	// c <- 1 // 注释掉，引发 timeout。
	<-w
}
