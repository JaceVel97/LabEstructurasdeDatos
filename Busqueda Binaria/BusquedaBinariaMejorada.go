package main

import "fmt"

func main() {
	var List = []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	end:=len(List) - 1

	response:= BusquedaBinariaRecursiva(List, 37, 0, end)

	if response {
		fmt.Println("Valor encontrado")
	} else {
		fmt.Println("Valor no encontrado")
	}
}

func BusquedaBinariaRecursiva (List[]int,  value int,  begin int, end int) bool {
	fmt.Println("Iteración")
	var middle int;
	if begin > end {
		return false
	} else {
		middle = (begin+end)/2
		if List[middle] < value {
			return BusquedaBinariaRecursiva(List, value, middle+1,end)
		} else if List[middle] > value {
			return BusquedaBinariaRecursiva(List, value, begin,middle-1)
		} else {
			return true
		}
	}
}