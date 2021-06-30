package main

import (
	"log"
	"os"
)

func main() {
	// 解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("echo.go")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}
}
