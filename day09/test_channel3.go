package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	close(done)
}

// 若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。
// 因此在这个例子中，用close(c)关闭管道代替done <- false依然能保证该程序产生相同的行为。
func main() {
	go aGoroutine()
	<-done
	println(msg)
}
