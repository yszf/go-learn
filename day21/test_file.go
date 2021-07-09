package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err) // remove astaxie: The directory is not empty.
	}
	os.RemoveAll("astaxie")
}
