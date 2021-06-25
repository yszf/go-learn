package main

import (
	"unsafe"

	"./my_buffer"
)

//#include <stdio.h>
import "C"

func main() {
	buf := my_buffer.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
