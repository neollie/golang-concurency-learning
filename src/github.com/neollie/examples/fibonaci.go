package examples

import (
	)

/*
Learn: channel of size one, one goroutine
*/

func getFibNum(i int, ch chan int) {
	defer close(ch)
	x, y := 0, 1
	for j := 0; j < i; j++ {
		x, y = y, x+y
	}
	ch <- x
}

// Return i-th fibonaci number ( at 0-th position we start with 0 ) with chanel of length one
func FibonaciNth(i int) int {
	ch := make(chan int, 1)
	go getFibNum(i, ch)
	return <-ch
}

/*
Learn : buffered channel, close channel, and range over channel
*/

// Generate fibonaci sequence of length n to channel ch
func genFibSeq(n int, ch chan int) {
	defer close(ch) // to inform receiver that routine ended
	// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x // send
		x, y = y, x+y
	}
}

// Get fibonaci sequence of length n with producer-consumer approach
func FibonaciSequence(n int) []int {
	ch := make(chan int, 3) // buffered channel with size equals to 10
	go genFibSeq(n, ch)
	seq := make([]int, 0, n)
	for v := range ch { // The loop for i := range c receives values from the channel repeatedly until it is closed. see explicit close in genFibSeq. Otherwise deadlock
		seq = append(seq, v)
	}
	return seq
}

/*
Learn 'select' statement
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Producer produces fib number to channel 'ch' in infinite loop and waits for stop message from channel 'quit'. Stop message is boolean 'true' but we dont care about value ( see row with code 'case <-quit')
Consumer consumes input from channel1 and send stop message to channel2 to stop producer
*/

func produceFibonaciNum(ch chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x: // TODO !
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func FibonaciNth2(i int) int {
	ch, quit := make(chan int), make(chan int)
	x := 0
	go func() {
		for j := 0; j <= i; j++ {
			x = <-ch
		}
		quit <- 0
	}() // see here has to be function invocation !!!
	produceFibonaciNum(ch, quit)
	return x
}
