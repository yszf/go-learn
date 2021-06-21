package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		fmt.Println("o.done = 0")
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

var (
	instance *singleton
	once     Once
)

func Instance(wg *sync.WaitGroup) *singleton {
	defer wg.Done()

	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go Instance(&wg)
	go Instance(&wg)
	go Instance(&wg)
	go Instance(&wg)

	wg.Wait()

}
