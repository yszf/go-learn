package main

import "C"

import (
	"fmt"
	"go-learn/day12/qsort2"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort2.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])),
		func(a, b unsafe.Pointer) int {
			pa, pb := (*int32)(a), (*int32)(b)
			return int(*pa - *pb)
		},
	)

	fmt.Println(values)
}
