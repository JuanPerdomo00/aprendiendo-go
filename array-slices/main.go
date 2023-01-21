package main

import "fmt"

func main() {
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println(array, len(array), cap(array))

	// slice
	slice := []int{0, 01, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos de slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:2])

}
