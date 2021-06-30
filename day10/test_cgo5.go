package main

// void SayHello2(char* s);
import "C"

import (
	"fmt"
)

//export SayHello
func SayHello2(s *C.char) {
	fmt.Print(C.GoString(s))
}

func main() {
	C.SayHello2(C.CString("Hello, World\n"))
}
