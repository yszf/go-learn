package main

import "fmt"

func main() {
	var d [0]int         // 定义一个长度为0的数组
	var e = [0]int{}     // 定义一个长度为0的数组
	var f = [...]int{}   // 定义一个长度为0的数组
	fmt.Println(d, e, f) // [] [] []

	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2
}
