package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
		fmt.Println("product : ", i*factor)
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println("Consume : ", v)
	}
}

func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出
	// time.Sleep(1 * time.Second)

	<-make(chan int)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
