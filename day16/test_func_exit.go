package main

//func main() {
//	defer func() {
//		select{}
//	}()
//}

// func main() {
// 	defer func() {
// 		for {

// 		}
// 	}()
// }

func main() {
	defer func() {
		<-make(chan bool) // fatal error: all goroutines are asleep - deadlock!
	}()

}
