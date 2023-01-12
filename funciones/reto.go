package funciones

import (
	"fmt"
	"math"
)

func areaRectangulo(largo int, ancho int) int {
	area := largo * ancho
	return area
}

func areaTrapecio(altura int, paraleloA int, paraleloB int) int {
	area := altura * ((paraleloA + paraleloB) / 2)
	return area
}

func areaCirculo(pi float64, radio float64) float64 {
	area := pi * math.Pow(radio, 2)
	return area

}

func main() {
	areaRec := areaRectangulo(20, 50)
	areaTra := areaTrapecio(100, 50, 70)
	areaCir := areaCirculo(math.Pi, 43.96)

	fmt.Println("El area del rectangulo es: ", areaRec)
	fmt.Println("El area del trapecio es: ", areaTra)
	fmt.Println("El area del circulo es: ", areaCir)
}
