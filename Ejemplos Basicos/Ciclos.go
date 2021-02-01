package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var j = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	/*
	for {
		fmt.Println("Infinito")
	}*/
}