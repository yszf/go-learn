package main

func test() (int, int) {
	return 1, 2
}

func main() {
	// s := make([]int, 2)
	// s = test() // Error: multiple-value test() in single-value context

	x, _ := test()
	println(x)
}
