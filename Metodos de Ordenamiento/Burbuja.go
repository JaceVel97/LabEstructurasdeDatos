package main

import "fmt"

func main(){
	var DisordelyList = []int{17, 7, 2, 98, 1, 23}
	var cubo = [][]int{{1,2},{4,5,6,5,4,8},{7,8,9,1}}
	var suma int =0
	fmt.Println("Tamaño: ", len(cubo))
	for i:=0;i<len(cubo);i++{
		suma += len(cubo[i])
	}
	fmt.Println("Tamaño total: ", suma)
	OrderedList := BubbleSort(DisordelyList)
	fmt.Println(OrderedList)
}

func BubbleSort(DisorderlyList []int) []int {
	var aux int
	for i := 0; i < len(DisorderlyList); i++ {
		for j := 0; j < len(DisorderlyList); j++ {
			if DisorderlyList[i] < DisorderlyList[j] {
				aux = DisorderlyList[i]
				DisorderlyList[i] = DisorderlyList[j]
				DisorderlyList[j] = aux
			}
		}
	}
	return DisorderlyList
}