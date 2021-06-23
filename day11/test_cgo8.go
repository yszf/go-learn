package main

/*
enum C {
    ONE = -8,
    TWO,
};
*/
import "C"
import "fmt"

func main() {
	var c C.enum_C = C.TWO
	fmt.Println(c)     // -7
	fmt.Println(C.ONE) // -8
	fmt.Println(C.TWO) // -7
}
