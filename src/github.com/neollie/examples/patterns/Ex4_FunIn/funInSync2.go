package examples

import (
	"github.com/neollie/examples/patterns/Ex3_TwoProducers"
)

// Fun In funs repeatedly one output from channel ch1 and ch2 (in FIFO style) into the channel ch
// Execution of producers is coupled (in one round we want ouptput from all producers ) but individual outputs are processed independently
// Regex [(s1 s2)|(s2 s1)]*

func funInSyncOnChan2(ch1 , ch2 chan string) (ch chan string) {
	ch = make(chan string)
	go func() {
		for {
			select {
			case s1 := <-ch1:
				ch <- s1
			case s2 := <-ch2:
				ch <- s2
			}
		}
	}()

	return ch
}


func RunAsync2(){
	ch:= funInSyncOnChan2(examples.Prod3("Moni"), examples.Prod3("Dia"))
	Cons4(ch)
}