package main

func add(x, y int) (z int) {
	z = x + y
	return
}

func sub(x, y int) (z int) {
	{ // 不能在⼀一个级别，引发 "z redeclared in this block" 错误。
		var z = x - y
		// return   // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}

func main() {
	println(add(1, 2))
	println(sub(10, 5))
}
