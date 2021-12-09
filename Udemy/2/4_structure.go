package main

import "fmt"


type People struct {
	Name string
	Age int
	Phone_no string
}

// add a function in man
func (p2 People) Get_age () int {
	return p2.Age;
}


func main() {

	var p People
	// p = People {"Sheiblu", 27, "016"} // Altanative
	p = People {
		Name : "Sheiblu",
		Age : 27, 
		Phone_no : "01671794064",
	}

	/* Altanative
	p := People {"Sheiblu", 27, "016"} 
	p := People { 
		Name : "Sheiblu",
		Age : 27, 
		Phone_no : "01671794064",
	} */

	p.Name ="Sheiblu Islam"
	fmt.Println(p)
	fmt.Println(p.Get_age())

}
