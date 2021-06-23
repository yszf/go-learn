package main

// const char* cs = "hello";
import "C"
import "day11/cgo_helper"

func main() {
	cgo_helper.PrintCString(C.cs) // cannot use C.cs (variable of type *_Ctype_char) as *cgo_helper._Ctype_char value in argument to cgo_helper.PrintCString
}
