package main

import "fmt"

func swap(v1, v2 *int) {
	temp := *v1
	*v1 = *v2
	*v2 = temp
}

func swap2(v1, v2 int) (int, int) {
	return v2 , v1 ;
}

func main() {
	m1, m2 := 20, 2
	pointer1 := &m2 // set m2 value position / pointer

	fmt.Println("m2:", m2, " pointer: ", pointer1)
	fmt.Println("m2:", m2, " &m1: ", &pointer1)
	fmt.Println("pointer1:", pointer1, " *pointer1: ", *pointer1)

	fmt.Println("Before swap m1: ", m1, " m2: ",m2)
	swap(&m1, &m2)
	fmt.Println("after swap m1: ", m1, " m2: ",m2)
	m3, m4 := swap2(m1, m2)
	fmt.Println("Again swap m1: ", m3, " m2: ",m4)
}
