package main

import (
	"fmt"
	"sync"
	"time"
)

func infinityCount(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		infinityCount("cat")
		wg.Done()
	}()
	wg.Wait()
}
