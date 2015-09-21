package examples

import "fmt"
import "time"

var i int = 0 // Probably bad, because 2 threads access it and ++ probably is not concurent or is it ?

// Generate int every t ms to chanel ch
// ret
func producer(t time.Duration, ch chan int) {
	for ; i < 100; i++ {
		time.Sleep(t)
		ch <- i
	}
}

func Example1() {
	ch := make(chan int)
	go producer(2000, ch)
	for {
		fmt.Println("waiting ... ", i)
		//if i < 100 { // Absolutely don't understand what does it do when it is uncommented...
			fmt.Println(<-ch) // Deadlock
		//}
	}
}
