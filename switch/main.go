package main

import "fmt"

func main() {

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 100

	switch {
	case value > 100:
		fmt.Println("value es mayor que 100")
	case value < 100:
		fmt.Println("value es menor que 100")

	default:
		fmt.Println("problrmas")
	}
}
