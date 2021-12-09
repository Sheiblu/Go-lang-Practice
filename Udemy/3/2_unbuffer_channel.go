package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//send
	go func() {
		c <- 1
		fmt.Println(1)
	}()

	//sniff
	val := <-c
	fmt.Println("val is : ", val)

	go func() {
		c <- 3
		fmt.Println(2)
	}()

	val = <-c
	fmt.Println("val is : ", val)
	fmt.Println("C is : ", c)
}
