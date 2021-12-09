package main

import (
	"fmt"
	"time"
)

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovery")
	}
}
func heavy() {
	i := 0
	defer cleanup()
	for {
		i++
		time.Sleep(time.Second * 1)
		fmt.Println("i is : ", i)
		if i == 5 {
			fmt.Println("Break")
			break
		}
	}
	// os.Exit(0)

}

func main() {
	// heavy();  // it's sync

	// But when use go it's async function but use 'Select {}'
	go heavy()
	select {}
}
