package main

import "fmt"

func main() {
	var arr []int
	arr2 := []int{1, 2, 3, 4}
	fmt.Println("Array2 : ", arr2)
	arr = append(arr2, 3)
	fmt.Println("Array1 : ", arr)
	fmt.Println("Array2 : ", arr2)

}
