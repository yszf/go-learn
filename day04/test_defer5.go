package main

func test() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, x, y) // y 闭包引⽤用 // defer: 10 20 120
	}(x) // x 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y) // x = 20 y = 120
}

func main() {
	test()
}
