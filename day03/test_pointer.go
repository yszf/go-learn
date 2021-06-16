package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type data struct {
		a int
	}

	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a) // 直接⽤用指针访问⺫⽬目标对象成员，⽆无须转换。

	//	x := 1234
	//	p2 := &x
	//	p2++ // 不能对指针做+-运算

	x := 0x12345678
	p3 := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p3)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
}
