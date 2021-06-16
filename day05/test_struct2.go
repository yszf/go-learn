package main

import "fmt"

func main() {
	type File struct {
		name string
		size int
		attr struct {
			perm  int
			owner int
		}
	}

	f := File{
		name: "test.txt",
		size: 1025,
		// attr: {0755, 1}, // Error: missing type in composite literal
	}
	f.attr.owner = 1
	f.attr.perm = 0755
	fmt.Println(f)
	var attr = struct {
		perm  int
		owner int
	}{2, 0755}
	f.attr = attr
	fmt.Println(f)
}
