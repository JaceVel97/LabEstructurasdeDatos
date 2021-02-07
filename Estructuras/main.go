package main

import (
	"fmt"
	"github.com/JaceVel97/Estructuras/ListaDoblementeEnlazada"
	"github.com/JaceVel97/Estructuras/ListaSimple"
	"github.com/JaceVel97/Estructuras/Pila"
)

func main()  {
	fmt.Println("Datos de Lista ")
	testList := ListaSimple.NewList()
	testList.InsertBack(1)
	testList.InsertBack(2)
	testList.InsertBack(3)
	testList.InsertBack(4)
	testList.InsertBack(5)
	fmt.Println("Tamaño de la lista: ", testList.GetSize())
	testList.ShowData()

	fmt.Println()
	fmt.Println("\nDatos de Pila ")
	testStack := Pila.NewStack()
	testStack.Push(3)
	testStack.Push(7)
	testStack.Push(11)

	fmt.Println("Valor en la cima de la pila ", testStack.Peek())
	fmt.Println("Tamaño de la pila: ", testStack.GetSize())
	fmt.Println("Se realiza un Pop a la pila y se obtiene ", testStack.Pop())
	fmt.Println("Tamaño de la pila: ", testStack.GetSize())
	fmt.Println("Se realiza un Pop a la pila y se obtiene ", testStack.Pop())
	fmt.Println("Tamaño de la pila: ", testStack.GetSize())
	fmt.Println("Valor en la cima de la pila ", testStack.Peek())
	fmt.Println("Tamaño de la pila: ", testStack.GetSize())


	//-------------------------------------------------------------------------------
	fmt.Println()
	fmt.Println()
	fmt.Println("Datos de Lista Doble")
	testListD := ListaDoblementeEnlazada.NewList()
	testListD.Insert(1)
	testListD.Insert(2)
	testListD.Insert(3)
	testListD.Insert(4)
	testListD.Insert(5)
	fmt.Println("Tamaño de la lista: ", testListD.GetSize())
	testListD.ShowData()
	fmt.Println()
	testListD.ShowDataBack()
}
