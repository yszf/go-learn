package main

func test(x, y int) {
	var z int = -1
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
		return
	}()
	println("x / y =", z)
}

func main() {
	test(2, 1) // x / y = 2
	test(2, 0) // x / y = 0
}
