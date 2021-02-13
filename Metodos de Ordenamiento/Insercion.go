package main

import "fmt"

func main(){
	var DisordelyList = []int{18, 8, 3, 99, 2, 24}
	OrderedList := InsertionSort(DisordelyList)
	fmt.Println(OrderedList)
}

func InsertionSort(DisorderlyList []int) []int {
	var aux int
	for i := 1; i < len(DisorderlyList); i++ {
		aux = DisorderlyList[i]
		for j := i - 1; j >= 0 && DisorderlyList[j] > aux; j-- {
			DisorderlyList[j+1] = DisorderlyList[j]
			DisorderlyList[j] = aux
		}
	}
	return DisorderlyList
}