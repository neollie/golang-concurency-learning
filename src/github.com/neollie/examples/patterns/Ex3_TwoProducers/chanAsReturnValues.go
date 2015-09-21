package examples

import (
	"log"
	"fmt"
)

const (
	Prod2Name = "Moni"
)

// Producer returns string channel where he produces sequence of strings infinitely
func Prod3(name string) (ch chan string) {
	ch = make(chan string)
	go func() {
		for x := 0;; x++ {
			ch <- fmt.Sprintf("prod2 - %s: %v", name, x)
		}
	}()
	return ch
}

// Consumes 5 numbers from both channels, log them and stop
func cons3(ch1 chan string, ch2 chan string) {
	for i:=0; i<5; i++ {
		x1 := <- ch1
		x2 := <- ch2
		log.Println(x1)
		log.Println(x2)
	}
}

func Run3() {
	ch1:=Prod3("Moni")
	ch2:=Prod3("Dia")
	cons3(ch1, ch2)
}


