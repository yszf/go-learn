package main

// 对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
var done = make(chan bool, 1)
var msg string

func aGoroutine() {
	msg = "hello, world"
	<-done
}

/*
	因为main线程中done <- true发送完成前，后台线程<-done接收已经开始，
	这保证msg = "hello, world"被执行了，所以之后println(msg)的msg已经被赋值过了。
	简而言之，后台线程首先对msg进行写入，然后从done中接收信号，随后main线程向done发送对应的信号，
	最后执行println函数完成。但是，若该Channel为带缓冲的
	（例如，done = make(chan bool, 1)），main线程的done <- true接收操作将不会被后台线程的<-done接收操作阻塞，
	该程序将无法保证打印出“hello, world”。

	对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。
	如果将C设置为0自然就对应无缓存的Channel，也即使第K个接收完成在第K个发送完成之前。
	因为无缓存的Channel只能同步发1个，也就简化为前面无缓存Channel的规则：对于从无缓冲Channel进行的接收，
	发生在对该Channel进行的发送完成之前。
*/
func main() {
	go aGoroutine()
	done <- true
	//done <- false
	println(msg)
}
