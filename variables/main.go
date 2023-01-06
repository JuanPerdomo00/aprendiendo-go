package main

import (
	"fmt"
)

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("PI con tipo de dato", pi)
	fmt.Println("PI sin tipo de dato", pi2)

	// Declaracion de variables
	base := 21
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// zero values

	var a int
	var b float64
	var name string
	var meama bool

	fmt.Println(a, b, name, meama)

	//calculo del area del cauadrado

	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * 2

	fmt.Println(areaCuadrado)

}
