package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	const (
		s = "abc"
		x
	)
	println(s, x)

	const (
		m byte = 100
		//		n int  = 1e20
	)

}
