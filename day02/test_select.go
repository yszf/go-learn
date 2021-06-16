package main

import (
	"fmt"
	"time"
)

func chann(ch chan int, stopCh chan bool) {
	var i int = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int = 2, 3
	select {
	case i1 = <-c1:
		fmt.Printf("receive ", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("send ", i2, "to c2\n")
	case i3, ok := <-c3:
		if ok {
			fmt.Printf("receive ", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("receive c :", c)
		case s := <-ch:
			fmt.Println("receive s :", s)
		case _ = <-stopCh:
			goto end
		}
	}

end:
	fmt.Println("end")
}
