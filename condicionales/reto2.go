/*
verificar si un usuario tiene la contrasena correcta
y retornar true, delo contrario retornar false
*/

package main

import (
	"fmt"
	"log"
)

type User struct {
	username string
	password string
}

func login(username string, password string) bool {
	if username == "jakepy" && password == "120malware" {
		return true
	}
	return false
}

func main() {
	username1 := User{"jakepy", "12malware"}

	if login(username1.username, username1.password) {
		fmt.Println("Bienvenido a la pagina")
	} else {
		log.Fatal("404 Not Found")
	}

}
