package main

import "C"

import (
	"fmt"

	_ "go-learn/day13/number"
)

func main() {
	println("Done")
}

//export goPrintln
func goPrintln(s *C.char) {
	fmt.Println("goPrintln:", C.GoString(s))
}
