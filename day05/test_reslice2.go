package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]                      // [2 3 4]
	fmt.Println(s1, len(s1), cap(s1)) // [2 3 4] 3 8
	s1[2] = 100
	s2 := s1[2:6]                     // [100 5 6 7]
	fmt.Println(s2, len(s2), cap(s2)) // [100 5 6 7] 4 6
	s2[3] = 200
	fmt.Println(s) // [0 1 2 3 100 5 6 200 8 9]
}
