package examples

import (
	"github.com/neollie/examples/patterns/Ex3_TwoProducers"
	"log"
)

// Fun In funs randomly inputs from channels into one output channel ch
// Execution of producers is decoupled
// Regex: [(s1|s2)]*

func funInAsyncOnChanells(ch1 , ch2 chan string) (ch chan string) {
	ch = make(chan string)
	go func() {	for { ch <- <-ch1 } }() //s1
	go func() { for { ch <- <-ch2 } }() //s2
	return ch
}

func Cons4(ch <-chan string) {
	for i:=0; i<10; i++ {
		log.Println(<-ch)
	}
}

func Run4(){
	ch:= funInAsyncOnChanells(examples.Prod3("Moni"), examples.Prod3("Dia"))
	Cons4(ch)
}