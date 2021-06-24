package main

//static void noreturn() {}
import "C"
import "fmt"

func main() {
	_, err := C.noreturn()
	fmt.Println(err) // nil

	v, _ := C.noreturn()
	fmt.Printf("%#v\n", v) // main._Ctype_void{}

	fmt.Println(C.noreturn()) // []
	// CGO生成的代码中，_Ctype_void类型对应一个0长的数组类型[0]byte，因此fmt.Println输出的是一个表示空数值的方括弧。
}
