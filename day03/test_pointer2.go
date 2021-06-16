package main

import (
	"fmt"
	"unsafe"
)

func test() *int {
	x := 100
	return &x // 在堆上分配 x 内存。但在内联时，也可能直接分配在目标栈。
}

func main() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset

	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2)        // Pointer -> *int
	*px = 200               // d.x = 200
	fmt.Printf("%#v\n", d)  // struct { s string; x int }{s:"abc", x:200}
}
