package main

import (
	"fmt"
	"time"
)

var now = time.Now()

func main() {
	fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
}

func init() {
	fmt.Printf("now: %v\n", now)
}

func init() {
	fmt.Printf("since: %v\n", time.Now().Sub(now))
}

func init() {
	fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
	w := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		w <- true
	}()
	<-w
}
