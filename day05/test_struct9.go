package main

import "fmt"

func main() {
	//不能同时嵌⼊入某一类型和其指针类型，因为它们名字相同
	type Resource struct {
		id int
	}
	type User struct {
		*Resource
		// Resource // Error: duplicate field Resource
		name string
	}
	u := User{
		&Resource{1},
		"Administrator",
	}
	println(u.id)
	println(u.Resource.id)
	fmt.Println(u) // {0xc000012088 Administrator}
}
