package main

func main() {
	// var i int
	// for {
	// 	println(i)
	// 	i++
	// 	if i > 2 {
	// 		goto BREAK
	// 	}
	// }
	goto BREAK
BREAK:
	println("break")
	//EXIT: // Error: label EXIT defined and not used

L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}
			print(x, ":", y, " ")
		}
		println()
	}

	x := 100
	switch {
	case x >= 0:
		if x == 100 {
			break
		}
		println(x)
	}
}
