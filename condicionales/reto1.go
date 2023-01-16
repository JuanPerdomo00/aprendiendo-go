/*
crear una funcion que duga su el numero
es par. Si es par, debe retornar true
delo contrario debe retornar false
*/

package main

import (
	"fmt"
	"log"
)

func isPar(n uint) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	var n uint = 5
	var valueBool bool = isPar(n)

	if valueBool {
		fmt.Printf("El numero %d es %t == numero par \n", n, valueBool)
	} else {
		log.Fatal("IMPAR")
	}

}
