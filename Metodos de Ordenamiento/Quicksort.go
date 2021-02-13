package main

import "fmt"

func main(){
	var DisordelyList = []int{19, 9, 4, 100, 3, 25}
	final := len(DisordelyList) - 1
	OrderedList := QuickSort(DisordelyList, 0, final)
	fmt.Println(OrderedList)
}

func QuickSort(DisordelyList []int, left int, right int) []int {
	pivote := DisordelyList[left] // tomamos primer elemento como pivote
	i := left                     // i realiza la búsqueda de izquierda a derecha
	j := right                    // j realiza la búsqueda de derecha a izquierda
	var aux int

	for i < j { // mientras no se crucen las búsquedas
		for DisordelyList[i] <= pivote && i < j {
			i++
		} // busca elemento mayor que pivote
		for DisordelyList[j] > pivote {
			j--
		} // busca elemento menor que pivote
		if i < j { // si no se han cruzado
			aux = DisordelyList[i] // los intercambia
			DisordelyList[i] = DisordelyList[j]
			DisordelyList[j] = aux
		}
	}
	DisordelyList[left] = DisordelyList[j] // se coloca el pivote en su lugar de forma que tendremos
	DisordelyList[j] = pivote              // los menores a su izquierda y los mayores a su derecha
	if left < j-1 {
		QuickSort(DisordelyList, left, j-1) // ordenamos subarray izquierdo
	}

	if j+1 < right {
		QuickSort(DisordelyList, j+1, right) // ordenamos subarray derecho
	}
	return DisordelyList
}