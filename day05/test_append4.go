package main

import "fmt"

func main() {
	s := make([]int, 0, 1)
	fmt.Println(s)
	fmt.Printf("&s = %p\n", &s)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		fmt.Println(s)
		fmt.Printf("&s = %p\n", &s)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}
