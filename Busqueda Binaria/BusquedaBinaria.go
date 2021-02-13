package main

import "fmt"

func main() {
	var List = []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}

	response:= Busqueda(List, 37)

	if response {
		fmt.Println("Valor encontrado")
	} else {
		fmt.Println("Valor no encontrado")
	}
}

func Busqueda(List []int, value int) bool{
	var  end, middle int

	counter:=0
	end = len(List) - 1

	for begin:=0; begin<=end; {
		fmt.Println("Iteración: ", counter)
		//Si la lista solo tiene un elemento
		if begin == end {
			if List[begin] == value {
				//Valor encontrado
				return true
			} else {
				//Valo no existe
				return false
			}
		}

		middle = (begin+end)/2 //Utilizado para dividir el array en dos

		if List[middle] == value {
			//Valor encontrado
			return true
		} else {
			// Si el numero buscado es menor
			if value < List[middle] {
				end = middle - 1 //Elige el elemneto de la derecha
			} else {
				begin = middle + 1 //Elige el elemento de la izquierda
			}
		}
		counter++
	}
	return false
}