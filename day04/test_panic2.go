package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover()) //defer panic
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func main() {
	test()
}
