package main

import "fmt"

type Data struct {
	x int
}

// 方法不过是一种特殊的函数，只需将其还原，就知道 receiver T 和 *T 的差别

func (self Data) ValueTest() { // func ValueTest(self Data)
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data)
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}
