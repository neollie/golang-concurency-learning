package examples

import "log"
import (
	"time"
	"fmt"
)

func prod(name string, sleepTime time.Duration) (ch chan string) {
	ch = make(chan string)
	go func() {
		for x := 0;; x++ {
			ch <- fmt.Sprintf("prod2 - %s: %v", name, x)
			time.Sleep(sleepTime)
		}
	}()
	return ch
}

func cons(ch chan string, timeout time.Duration) {
	for i:=0; i < 10; i++ {
		select {
		case s := <- ch: log.Println(s)
		case <-time.After(timeout): fmt.Printf("!!!Timeout for phase %v\n", i); return; // See return here
		}
	}
}



func RunTimeout() {
	sleepTime:= time.Duration(5)
	timeout:=time.Duration(3)
	cons(prod("Moni",sleepTime), timeout)
}
