package main

func add(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func main() {
	println(add(1, 2)) // 输出: 103
}
