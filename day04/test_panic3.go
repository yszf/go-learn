package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover()) //test panic
		//println(recover().(string))
	}()

	defer fmt.Println(recover()) // nil

	defer func() {
		//	recover()
		func() {
			println("defer inner")
			recover() // 无效！
		}()

	}()
	panic("test panic") // panic: test panic, 未捕获
}

func main() {
	test()
}
