package main

import (
	"go-learn/day13/my_buffer"
	"unsafe"
)

//#include <stdio.h>
import "C"

func main() {
	buf := my_buffer.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
