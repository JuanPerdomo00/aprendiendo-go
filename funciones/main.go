package main

import "fmt"

func normalFunc(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Printf("%d %d %s \n", a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunc("Mensaje secreto")
	tripeArgument(12, 14, "hola")

	value := returnValue(3)
	fmt.Printf("%d \n", value)

	value2, value3 := doubleReturn(2)
	fmt.Printf("value2: %d, value3: %d \n", value2, value3)

	// descartar una variable con el _
	value4, _ := doubleReturn(2)
	fmt.Printf("value4: %d \n", value4)
}
