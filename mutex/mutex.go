package main

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	Age int64
}

func addAge(cl *Client, add int64, mu *sync.Mutex) {
	mu.Lock()
	cl.Age = cl.Age + add
	mu.Unlock()
}

func main() {
	cl := &Client{}
	mu := &sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go addAge(cl, 1, mu)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("%#v\n", cl)
}
