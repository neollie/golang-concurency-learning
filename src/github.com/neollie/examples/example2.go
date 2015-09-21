package examples

import (
	"fmt"
	"math/rand"
	//"time"
)

var (
	goroutineCount = 4
	c = make(chan int)
)

func goroutine(name string) {
	for i:=0; i < 20; i++ {
		fmt.Printf("Goroutine %s : i = %d\n", name, i)
	}
	c <- 1
}

// generate random string of length n
func genRandStr(n int) (str string, err error) {
	if (n < 0) {
		return "", nil //TODO
	}

	letters := "abcdefghijklmnopqrstuvwyzx"
	l:=letters[:]
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		pos:= rand.Intn(len(letters))
		s[i] = l[pos]
	}
	return string(s), nil
}

func Example2() {
	for i:=0; i < goroutineCount; i++ {
		name,_ := genRandStr(6)
		// fmt.Println(name)
		go goroutine(name)
	}

	for i:=0; i < goroutineCount; i++ {
		<- c
	}
	//time.Sleep(10*time.Second)
	fmt.Println("Main end")
}