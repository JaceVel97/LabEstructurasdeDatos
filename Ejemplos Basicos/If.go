package main

import "fmt"

func main() {
	encendido := false
	var edad uint8 = 18

	//If básico
	if encendido {
		fmt.Println("Está encendido")
	}

	//If - Else
	if encendido {
		fmt.Println("Está encendido")
	} else {
		fmt.Println("Está apagado")
	}

	//If, Else if
	if edad > 150 {
		fmt.Println("¿Eres inmortal?")
	} else if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad >= 15 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	//If con declaración corta

	if edad2 := 25; edad2 > 150 {
		fmt.Println("¿Eres inmortal?")
	} else if edad2 >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad2 < 18 && edad2 > 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Edad fuera del rango")
	}


}
