package main

/*
extern char* NewGoString(char* );
extern void FreeGoString(char* );
extern void PrintGoString(char* );

static void printString(const char* s) {
    char* gs = NewGoString((char*)s);
    PrintGoString(gs);
    FreeGoString(gs);
}
*/
import "C"

import (
	"go-learn/day12/object"
	"unsafe"
)

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := object.NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := object.ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(p *C.char) {
	id := object.ObjectId(uintptr(unsafe.Pointer(p)))
	gs := id.Get().(string)
	println(gs)
}

func main() {
	C.printString(C.CString("hello"))
}
