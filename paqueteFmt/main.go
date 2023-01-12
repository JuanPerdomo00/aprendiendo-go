package main

import "fmt"

func main() {
	// Declaracion
	helloMessage := "Hello"
	worldMessage := "World"

	//! Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//! Printf
	nombre := "lalo"
	apellido := "lelo"
	fmt.Printf("%s %s bienvenido a este programa \n", nombre, apellido)
	// %v cuando no tienes claro que tipo de dato va
	fmt.Printf("%v %v bienvenido a este programa \n", nombre, apellido)

	// Sprintf
	/*
	*bool:                    %t
	*int, int8 etc.:          %d
	*uint, uint8 etc.:        %d, %#x if printed with %#v
	*float32, complex64, etc: %g
	*string:                  %s
	*chan:                    %p
	*pointer:                 %p
	 */

	message := fmt.Sprintf("%s %s bienvenido a este programa \n", nombre, apellido)
	fmt.Println(message)

	// Tipo de datos ver con %T
	var comida string = "pizza"
	var numeroFav int = -77
	var numeroFavP uint = 77
	var numeroPi float64 = 3.141516

	fmt.Println("===============================")
	fmt.Printf("comida: %T\n", comida)
	fmt.Println("===============================")
	fmt.Printf("numeroFav: %T\n", numeroFav)
	fmt.Println("===============================")
	fmt.Printf("numeroFavP: %T\n", numeroFavP)
	fmt.Println("===============================")
	fmt.Printf("numeroPi: %T\n", numeroPi)
	fmt.Println("===============================")
}
