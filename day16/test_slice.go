package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// 切片会导致整个底层数组被锁定，底层数组无法释放内存。如果底层数组较大会对内存产生很大的压力
	headerMap := make(map[string][]byte)

	for i := 0; i < 5; i++ {
		name := "echo.go"
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		headerMap[name] = data[:1]
	}

	// do some thing
}
