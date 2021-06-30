package main

import "C"

import (
	"fmt"
	"go-learn/day12/qsort3"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort3.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
