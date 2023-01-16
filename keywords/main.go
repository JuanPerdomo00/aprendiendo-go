package main

/*
palabras claves

defer -> para ejecutar una funcion caundo el programa muera
continue -> se salta la iteracion en caso de expecificarla
break -> detiene el ciclo
*/

import "fmt"

func saludar() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func main() {
	defer saludar()

	// break y continue

	for i := 0; i < 10; i += 1 {
		if i == 5 {
			continue
		}

		fmt.Println(i)
		if i == 8 {
			break
		}
	}

}
