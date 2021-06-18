package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// type StringHeader struct {
	// 	Data uintptr
	// 	Len  int
	// }

	s := "hello, world"
	// var data = [...]byte{
	// 	'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',
	// }

	hello := s[:5]
	world := s[7:]

	fmt.Println(hello) // hello
	fmt.Println(world) // world

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

	fmt.Println("-----------------------------")
	fmt.Printf("%#v\n", []byte("Hello, 世界")) // []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c} // utf-8 encoding

	fmt.Println("-----------------------------")
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �  界abc

	fmt.Println("-----------------------------")
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
	// 0 65533  // \uFFFD, 对应 �
	// 1 0      // 空字符
	// 2 0      // 空字符
	// 3 30028  // 界
	// 6 97     // a
	// 7 98     // b
	// 8 99     // c

	fmt.Println("-----------------------------")
	for i, c := range "世界abc" {
		fmt.Println(i, c)
	}

	fmt.Println("------------------------------")
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	fmt.Println("------------------------------")
	const s3 = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%d %x\n", i, s3[i])
	}

}
