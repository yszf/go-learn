package main

import "fmt"

type Node struct {
	_    int
	id   int
	data *byte
	next *Node
}

func main() {
	n1 := Node{
		id:   1,
		data: nil,
	}

	n2 := Node{
		id:   2,
		data: nil,
		next: &n1,
	}

	fmt.Println(n1) // {0 1 <nil> <nil>}
	fmt.Println(n2) // {0 2 <nil> 0xc0000a63a0}

	type User struct {
		name string
		age  int
	}

	u1 := User{"Tom", 20}
	//	u2 := User{"Bob"} // too few values in User{...}
	u2 := User{name: "Bob"}
	fmt.Println(u1) // {Tom 20}
	fmt.Println(u2) // {Bob 0}

}
