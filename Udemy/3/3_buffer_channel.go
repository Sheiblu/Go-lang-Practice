package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func main() {
	c := make(chan int, 3)
	//send
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
		fmt.Println(1)
	}()

	//sniff
	for i := range c {
		fmt.Println("val is : ", i)
	}

	car := make(chan Car, 3)
	go func() {
		car <- Car{"Bus"}
		car <- Car{"Car"}
		car <- Car{"CNG"}
		close(car)
		fmt.Println(2)
	}()

	for i := range car {
		fmt.Println("Car val is : ", i.Name)
	}
}
