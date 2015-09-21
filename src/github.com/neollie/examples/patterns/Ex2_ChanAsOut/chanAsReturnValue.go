package examples

import (
	"log"
	"fmt"
)

const (
	Prod2Name = "Moni"
)
// 1. Example : producer and consumer communicate with channel provided by producer

// Produce return channel where he produces sequence of string infinitely
func prod2(name string) (ch chan string) {
	ch = make(chan string)
	go func() {
		for x := 0;; x++ {
			ch <- fmt.Sprintf("prod2 - %s: %v", name, x)
		}
	}()
	return ch
}
// Consume 5 numbers from the prod2's channel, log them and stop
func cons2(ch chan string) {
	for i:=0; i<5; i++ {
		x := <- ch
		log.Println(x)
	}
}

func Run2() {
	ch:=prod2(Prod2Name)
	cons2(ch)
}

