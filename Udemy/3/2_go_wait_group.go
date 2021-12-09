package main

import (
	"fmt"
	"sync"
)

func main() {
	//Wait group
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done")
}
