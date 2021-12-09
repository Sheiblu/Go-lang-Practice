package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func heavy() {
	i := 0
	defer wg.Done()
	for {
		i++
		time.Sleep(time.Second * 1)
		fmt.Println("i is : ", i)
		if i == 5 {
			fmt.Println("Break")
			break
		}
	}
	fmt.Println("done")
}

func main() {
	// heavy();  // it's sync
	wg.Add(1)
	go heavy()
	wg.Wait()
}
