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
	//la variable graph me ayudara a guardar toda el codigo del grafico.
	var graph string = "digraph List {\n" //Este es el encabezado no debe cambiar nada solo se puede cambiar el nombre List por el de 
					     // de su preferencia lo demás se queda así.
	graph += "rankdir=LR;" //Esto es solo para que la grafica se ordene en modo horizontal, puede cambiar si es necesario si se quiere
				//vertical se cambia LR por TB.
	//Esta linea es para modificar como se ve el nodo tanto el color interno como los bordes.
	graph += "node [shape = record, color=blue , style=filled, fillcolor=skyblue];"
	var nodes string = "" //Manejara todos los nodos la declaracion
	var pointers string = "" //Manejara todos los punteros y conexiones, es mejor separarlo para que no haya conflicto luego.
	for auxiliar != nil {
		//Como los nodos deben tener un nombre unico entonces le concatene su valor, entonces si un nodo tiene dentro un 5 entonces
		//se llamaria node5 y ahora bien el label es lo que tendra dentro del nodo aqui puede ir el nombre de la tienda.
		nodes += "Node" + strconv.Itoa(auxiliar.value) + "[label=\"" + strconv.Itoa(auxiliar.value) + "\"]\n"
		if auxiliar.next != nil{
			//Aqui se almacenan los punteros permite apuntar del actual al siguiente
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