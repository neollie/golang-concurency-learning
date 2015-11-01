package examples

import (
	"github.com/neollie/examples/patterns/Ex3_TwoProducers"
)

// Fun In funs repetedly one output from channel ch1 followed by output from ch2 into the channel ch
// Execution of producers is coupled (we want from each one output than continue)
// Thus we ALWAYS BLOCK ch2 until ch1 is not ready (see. funInSyncOnChanNoBlock for enhancement)
// Regex: (s1 s2)*

func funInSyncOnChan(ch1 , ch2 chan string) (ch chan string) {
	ch = make(chan string)
	go func() {
		for {
			ch <- <-ch1 // sends string s1
			ch <- <-ch2 //sends string s2
		}
	}()

	return ch
}


func RunAsync(){
	ch:= funInSyncOnChan(examples.Prod3("Moni"), examples.Prod3("Dia"))
	Cons4(ch)
}
