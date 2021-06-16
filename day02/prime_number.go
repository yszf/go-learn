package main

import "fmt"

func PrintPrimeNumber(num int) {
	fmt.Println("prime number in", num)

	for i := 2; i < num; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if 0 == i%j {
				flag = false
				break
			}
		}

		if flag {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func main() {
	PrintPrimeNumber(100)
}
