package main

// 可以将 channel 隐式转换为单向队列，只收或只发
func main() {
	c := make(chan int, 3)
	var send chan<- int = c // send-only
	var recv <-chan int = c // receive-only
	send <- 1
	// <-send // Error: receive from send-only type chan<- int
	<-recv
	// recv <- 2 // Error: send to receive-only type <-chan int

	// 不能将单向 channel 转换为普通 channel
	//	d := (chan int)(send) // Error: cannot convert type chan<- int to type chan int
	//	e := (chan int)(recv) // Error: cannot convert type <-chan int to type chan int
}
