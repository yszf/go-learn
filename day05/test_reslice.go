package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[2:5]
	fmt.Println(s1, len(s1), cap(s1)) // [2 3 4] 3 8

	s2 := s1[2:6:7]
	fmt.Println(s2, len(s2), cap(s2)) // [4 5 6 7] 4 5

	s3 := s2[3:6] // slice bounds out of range
	fmt.Println(s3, len(s3), cap(s3))
}
