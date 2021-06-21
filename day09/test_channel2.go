package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	done <- false
}

// Channel通信是在Goroutine之间进行同步的主要方法。
// 在无缓存的Channel上的每一次发送操作都有与其对应的接收操作相配对，
// 发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine上执行2个操作很容易导致死锁）。
// 无缓存的Channel上的发送操作总在对应的接收操作完成前发生.
func main() {
	go aGoroutine()
	<-done
	println(msg)
}
