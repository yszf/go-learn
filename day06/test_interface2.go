package main

import "fmt"

//空接口 interface{} 没有任何方法签名，也就意味着任何类型都实现了空接口。其作用类似面向对象语⾔言中的根对象 object。
func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

func main() {
	Print(1)               // int: 1
	Print("Hello, World!") // string: Hello, World!
}
