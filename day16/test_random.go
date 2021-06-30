package main

import "fmt"

// 随机数的一个特点是不好预测。如果一个随机数的输出是可以简单预测的，那么一般会称为伪随机数。
// 基于select语言特性构造的随机数生成器。
func main() {
	for i := range random(100) {
		fmt.Println(i)
	}
}

func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}
