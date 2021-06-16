package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	path, err := exec.LookPath(os.Args[0])
	if nil != err {
		fmt.Println(err)
		return
	}

	res, err := filepath.Abs(path)
	if nil == err {
		fmt.Println(res)
	}
}
