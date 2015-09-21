package examples

import (
	"log"
	"fmt"
)

const (
	Prod1Name = "Moni"
)
// 1. Example : producer and consumer communicate with shared channel or channel created by consumer

// Produce increasing sequence of numbers to the provided channel starting from 0 and incremented by one
func Prod1(name string, ch chan string) {
	for x:=0;;x++ {
		ch <- fmt.Sprintf("prod2 - %s: %v", name, x)
	}
}
// Consume 5 number from the channel, log them and stop
func Cons1() {
	ch := make(chan string) // Channel created by consumer
	go Prod1(Prod1Name, ch)
	for i:=0; i<5; i++ {
		str := <- ch
		log.Println(str)
	}
}


