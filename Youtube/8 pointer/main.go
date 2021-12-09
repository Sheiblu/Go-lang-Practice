package main

import "fmt"

type Car struct {
	Top_speed int
	Gas_padel int
}

func (c Car) khm() int {
	c.Top_speed = 10
	return c.Top_speed * c.Gas_padel
}

func (c *Car) mhm() int {
	c.Top_speed = 33
	return c.Top_speed * c.Gas_padel
}

func main() {
	fmt.Println("Sheiblu")
	c := Car{5, 5}
	fmt.Println(c.Top_speed)
	fmt.Println(c.Gas_padel)

	fmt.Println(c.khm())

	fmt.Println(c.Top_speed)
	fmt.Println(c.Gas_padel)

	fmt.Println(c.mhm())

	fmt.Println(c.Top_speed)
	fmt.Println(c.Gas_padel)
}
