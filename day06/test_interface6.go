package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 接口对象由接口表 (interface table) 指针和数据指针组成。
// 只有 tab 和 data 都为 nil 时，接口才等于 nil。
func main() {
	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil
	type iface struct {
		itab, data uintptr
	}

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)                             // true {0 0}
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil()) // false {6155104 0} true
}
