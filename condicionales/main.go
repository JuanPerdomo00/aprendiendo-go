package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Si es 1")
	} else {
		fmt.Println("No es uno")
	}

	// with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad con AND &&")
	}

	//with or
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("SI me cumplo con OR ||")
	}

	// convertir de texto a numero
	// value, err := strconv.Atoi("31")
	value, err := strconv.Atoi("ww")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value: %d\n", value)
}
