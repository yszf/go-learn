package main

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {

}

func main() {

	c := Black
	test(c)

	x := 1
	//	test(x)
	test(Color(x))

	test(1)

}
