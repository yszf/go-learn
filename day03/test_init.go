package main

import "fmt"

func main() {

	var a struct{ x int } = struct{ x int }{100}
	var b []int = []int{1, 2, 3}
	c := struct {
		x int
		y string
	}{1, "hello"}
	// var a = struct{ x int }{100}
	// var b = []int{1, 2, 3}

	println(a.x, c.x, c.y)

	for _, val := range b {
		fmt.Printf("%d ", val)
	}
}
