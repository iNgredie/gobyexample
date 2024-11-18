package main

import (
	"fmt"
)

func readChan(ch chan int) {
	value := <-ch
	fmt.Println("CHAN VALUE ", value)
}

func main() {
	fmt.Println("START MAIN")
	var ch chan int

	ch = make(chan int, 1)
	ch <- 100
	go readChan(ch)

	//time.Sleep(1 * time.Second)
	ch <- 200

	fmt.Println("END MAIN")
}
