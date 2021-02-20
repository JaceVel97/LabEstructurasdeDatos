package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type node struct {
	next *node
	value int
}

type list struct {
	first *node
}

func NewList() *list {
	return &list{nil}
}

//Equal list
func (elist *list) Insert(value int) {
	newNode := &node{nil, -1}
	newNode.value = value
	if elist.first == nil {
		elist.first = newNode
	} else {
		auxiliar := elist.first
		for elist.first.next != nil {
			elist.first = elist.first.next
		}
		elist.first.next = newNode
		elist.first = auxiliar
	}
}

func (elist *list)InsertBack(value int) {
	newNode := &node{nil, -1}
	newNode.value = value
	if elist.first == nil {
		elist.first = newNode
	} else {
		newNode.next = elist.first
		elist.first = newNode
	}
}

func (elist *list)ShowData() {
	auxiliar := elist.first
	for auxiliar != nil {
		fmt.Println("Valor ", auxiliar.value)
		auxiliar = auxiliar.next
	}
}

func (elist *list)GraphData() {
	auxiliar := elist.first
	var graph string = "digraph List {\n"
	graph += "rankdir=LR;"
	graph += "node [shape = record, color=blue , style=filled, fillcolor=skyblue];"
	var nodes string = ""
	var pointers string = ""
	for auxiliar != nil {
		nodes += "Node" + strconv.Itoa(auxiliar.value) + "[label=\"" + strconv.Itoa(auxiliar.value) + "\"]\n"
		if auxiliar.next != nil{
			pointers += "Node" + strconv.Itoa(auxiliar.value) + "->Node" + strconv.Itoa(auxiliar.next.value) + ";\n"
		}
		auxiliar = auxiliar.next
	}
	graph += nodes + "\n" + pointers
	graph += "\n}"
	//Visualizar código de Graphviz, es opcional solo para verlo
	fmt.Println(graph)
	//Creacion del archivo .dot
	data := []byte(graph)	//Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen
	if err != nil {
		log.Fatal(err)
	}
	//Creación de la imagen
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
										// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777) //Se mantiene igual
	ioutil.WriteFile("List.png", cmd, os.FileMode(mode)) //Creacion de la imagen
}

func (elist *list)GetSize() int{
	auxiliar := elist.first
	var counter int = 0
	for auxiliar != nil {
		counter++
		auxiliar = auxiliar.next
	}
	return counter
}


func main(){
	testList := NewList()
	testList.Insert(1)
	testList.Insert(2)
	testList.Insert(3)
	testList.Insert(4)
	testList.Insert(5)
	testList.ShowData() //Lista en consola
	testList.GraphData() //Grafica de la lista
}