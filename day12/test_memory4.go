package main

/*
#include <stdio.h>
extern int* getGoPtr();

static void Main() {
    int* p = getGoPtr(); // panic: runtime error: cgo result has Go pointer
    *p = 42;
	printf("*p = %d\n", *p);
}
*/
import "C"

// go tool cgo test_memory4.go
// 设置环境变量GODEBUG=cgocheck=0来关闭指针检查行为。
// GODEBUG=cgocheck=0 go run test_memory4.go
func main() {
	C.Main()
}

//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}
