package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize) {
    return malloc(memsize);
}
*/
import "C"
import "unsafe"

// C语言的内存在分配之后就是稳定的，但是Go语言因为函数栈的动态伸缩可能导致栈中内存地址的移动(这是Go和C内存模型的最大差异)。如果C语言持有的是移动之前的Go指针，那么以旧指针访问Go对象时会导致程序崩溃。

// 因为Go语言实现的限制，我们无法在Go语言中创建大于2GB内存的切片
func makeByteSlize(n int) []byte {
	p := C.makeslice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n] // panic: runtime error: slice bounds out of range [::2147483649] with length 2147483648
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	s := makeByteSlize(1<<31 + 1)
	s[len(s)-1] = 255
	println(s[len(s)-1])
	freeByteSlice(s)
}
