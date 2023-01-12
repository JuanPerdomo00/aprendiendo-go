package main

import (
	"fmt"
)

func main() {
	// For condicional
	for i := 0; i <= 10; i += 1 {
		fmt.Printf("%d\n", i)
	}

	fmt.Println()
	// For while
	counter := 0
	for counter < 10 {
		fmt.Printf("%d \n", counter)
		counter += 1
	}

	// For forever
	counterForever := 0
	for {
		fmt.Printf("%d \n", counterForever)
		counterForever += 9
		break
	}

	//For Range
	arreglo := [8]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arreglo:", arreglo)

	fmt.Println("Primer ejemplo")
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}

	fmt.Println("Segundo ejemplo")
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}

	fmt.Println("Tercer ejemplo")
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}

	count := 10
	for count > 0 {
		fmt.Printf("%d\n", count)
		count -= 1
	}

}
