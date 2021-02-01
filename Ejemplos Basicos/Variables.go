package main

import "fmt"

func main() {
	var estatura float32 = 1.72 // Asignando valor al declarar
	var edad = 23               // Infiriendo tipo de dato de la asignación
	var adulto bool             // Declaración tradicional
	adulto = true               // Posterior asignación de valor
	nombre := "José Véliz"      // Declaración corta

	fmt.Println("-->Uso de variables<--")
	fmt.Println("Edad: ", edad)
	fmt.Println("Estatura: ", estatura)
	fmt.Println("Nombre: ", nombre)
	fmt.Println("¿Es adulto? ", adulto)
}