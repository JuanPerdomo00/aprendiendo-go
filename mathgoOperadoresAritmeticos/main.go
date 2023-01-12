package main

import "fmt"

func main() {
	x := 10
	y := 50

	// suma
	result := x + y

	fmt.Println("la suma es: ", result)

	// resta
	result = y - x
	fmt.Println("la resta es: ", result)

	// multiplicacion
	result = x * y
	fmt.Println("la multiplicacion es: ", result)

	// divicion
	result = y / x
	fmt.Println("la divicion es: ", result)

	// modulo
	result = y % x
	fmt.Println("modulo: ", result)

	// incremental
	// x += 1
	x++
	fmt.Println(x)

	// decremental
	// x -= 1
	x--
	fmt.Println(x)

}
