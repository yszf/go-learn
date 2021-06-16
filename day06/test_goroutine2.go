package main

import (
	"runtime"
	"sync"
)

// 调用 runtime.Goexit 将立即终止当前 goroutine 执行，调度器确保所有已注册 defer延迟调用被执行
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer println("A.defer")
		func() {
			defer println("B.defer")
			runtime.Goexit() // 终⽌止当前 goroutine
			println("B")     // 不会执⾏行
		}()
		println("A") // 不会执⾏行
	}()
	wg.Wait()
}
