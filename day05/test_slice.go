package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[1:4:5] // [low:high:max]

	fmt.Println(data)
	fmt.Println(slice)
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))

	fmt.Println(data[:6:8], len(data[:6:8]), cap(data[:6:8])) // 省略low

	fmt.Println(data[5:], len(data[5:]), cap(data[5:])) // 省略high、max

	fmt.Println(data[:3], len(data[:3]), cap(data[:3])) // 省略low、max

	fmt.Println(data[:], len(data[:]), cap(data[:])) // 省略low、high、max

}
