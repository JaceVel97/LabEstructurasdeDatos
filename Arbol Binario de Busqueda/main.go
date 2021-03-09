package main

import (
	abb "Trees/ABB"
	"fmt"
)

func main() {
	TestTreeAbb := abb.NewTree()
	TestTreeAbb.Insert(8)
	TestTreeAbb.Insert(4)
	TestTreeAbb.Insert(12)
	TestTreeAbb.Insert(6)
	TestTreeAbb.Insert(2)
	TestTreeAbb.Insert(10)
	TestTreeAbb.Insert(14)
	TestTreeAbb.GetOrder(0)
	fmt.Println("-------------")
	TestTreeAbb.GetOrder(1)
	fmt.Println("-------------")
	TestTreeAbb.GetOrder(2)
	fmt.Println("-------------")
	fmt.Println(TestTreeAbb.BuscarNode(12))
	fmt.Println(TestTreeAbb.BuscarBool(15))
}