package main

import (
	"fmt"
	"time"
)

func countWithChannel(
	thing string,
	c chan string,
) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	c := make(chan string)
	go countWithChannel("dog", c)

	for msg := range c {
		fmt.Println(msg)
	}
}
