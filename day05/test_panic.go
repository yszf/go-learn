package main

import "fmt"

func except() {
	fmt.Println(recover())
}

func test() {
	defer except()
	panic("test panic")
}

func main() {
	test()
}
