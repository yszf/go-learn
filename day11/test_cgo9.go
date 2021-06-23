package main

/*
#include <string.h>
char arr[10];
char *s = "Hello";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// 如果不希望单独分配内存，可以在Go语言中直接访问C语言的内存空间
func main() {
	// 通过 reflect.SliceHeader 转换
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10
	fmt.Println(arr0Hdr) // &{4159984 10 10}

	// 通过切片语法转换
	arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
	fmt.Println(arr1) // [0 0 0 0 0 0 0 0 0 0]

	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
	s0Hdr.Len = int(C.strlen(C.s))
	fmt.Println(s0Hdr) // &{3846864 5}

	sLen := int(C.strlen(C.s))
	fmt.Println((*[31]byte)(unsafe.Pointer(C.s))[:sLen:sLen]) // [72 101 108 108 111]
	s1 := string((*[31]byte)(unsafe.Pointer(C.s))[:sLen:sLen])
	fmt.Println(s1) // Hello
}
