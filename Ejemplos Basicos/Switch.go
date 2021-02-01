package main

import "fmt"

func main() {
	var edad uint8 = 18

	//Switch con declaración corta
	switch dias := 1; dias {
	case 0:
		fmt.Println("Lunes")
	case 1:
		fmt.Println("Martes")
	case 2:
		fmt.Println("Miercoles")
	case 3:
		fmt.Println("Jueves")
	case 4:
		fmt.Println("Viernes")
	case 5:
		fmt.Println("Sabado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Desconocido")
	}

	switch {
	case edad >= 150:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Eres menor de edad")
	}
}