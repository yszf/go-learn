package main

// void PrintCString(char*);
// #include <stdio.h>
// char* str = "hello"; // (static) undefined reference to `cs' // multiple definition of `str'
import "C"

func main() {
	C.puts(C.CString("Hello, World"))

	C.puts(C.str)

	C.PrintCString(C.CString("Hello, World"))

	//	C.PrintCString(C.str)
}

//export PrintCString
func PrintCString(cs *C.char) {
	C.puts(cs)
}
