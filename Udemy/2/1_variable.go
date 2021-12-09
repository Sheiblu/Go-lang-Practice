package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	name = "Sheiblu"
	fmt.Println(name)

	var age = 26
	fmt.Println("Age : ", age, " From BD.")

	var (
		num1 = 2
		num2 = 3
	)
	total := num1 + num2
	fmt.Println("Total sum: ", total)
	fmt.Print(strings.Contains("Sheiblu", "u"))
}
