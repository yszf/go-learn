package main

import (
	"fmt"
	"unsafe"
)

var x, y int

//g, h := 123, "hello"

func main() {
	fmt.Println("hello, world!")

	var stockcode int = 123
	var enddate string = "2021-2-22"
	var url string = "Code=%d&endDate=%s"

	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var b, c int = 1, 2
	var d bool
	var e string = "jack"
	fmt.Println(b, c, d, e) // 1 2 false ""
	fmt.Printf("%v %v %v %q\n", b, c, d, e)

	var p *int
	fmt.Println(p) // nil

	var p2 []int
	fmt.Println(p2) // []

	var p3 map[string]int
	fmt.Println(p3) // map[]

	var p4 chan int
	fmt.Println(p4) // nil

	var p5 func(string) int
	fmt.Println(p5) // nil

	var p6 error
	fmt.Println(p6) // nil

	var intVal int
	//	intVal := 1 // 没有申明新的变量
	intVal, intVal1 := 1, 2
	fmt.Println(intVal, intVal1)

	var (
		m int
		n bool
	)
	fmt.Println(m, n)

	const (
		x = "abc"
		y = len(x)
		z = unsafe.Sizeof(x)
	)
	fmt.Println(x, y, z)

	const (
		x2 = iota
		y2 //= iota
		z2 //= iota
	)
	fmt.Println(x2, y2, z2)

	const (
		i = 1 << iota
		j = 3 << iota
		k = 5 << iota
		l
		ll
	)
	fmt.Println(i, j, k, l, ll) // 1 6 20 40 80
}
