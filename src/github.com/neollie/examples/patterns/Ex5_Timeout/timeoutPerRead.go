package examples

import "log"
import (
	"time"
	"fmt"
)

// Producer produces infinitely string, after each one he sleeps second longer
func prod5(name string) (ch chan string) {
	ch = make(chan string)
	go func() {
		for i := 0;; i++ {
			ch <- fmt.Sprintf("prod2 - %s: %v", name, i)
			time.Sleep(time.Duration(i)*time.Second)
		}
	}()
	return ch
}

// Consumers wait limited time per read
// TODO - this is not good , there is no discard here , in next phase old string can come !!!
func cons5(ch chan string, timeout time.Duration) {
	for i:=0; i < 10; i++ {
		fmt.Printf("-Wait for %v-th string\n", i)
		select {
		case s := <- ch: log.Println(s)
		case <-time.After(timeout): fmt.Printf("!!!Timeout for phase %v\n", i); // No return here
		}
	}
}

func RunReadTimeout() {
	timeout:=time.Duration(5)*time.Second
	cons5(prod5("Moni"), timeout)
}
