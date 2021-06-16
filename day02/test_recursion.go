package main

import "fmt"

// 阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))

	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
