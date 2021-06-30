package main

import (
	"log"
	"os"
)

func main() {
	// defer在函数退出时才能执行，在for执行defer会导致资源延迟释放
	for i := 0; i < 5; i++ {
		f, err := os.Open("echo.go")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
}
