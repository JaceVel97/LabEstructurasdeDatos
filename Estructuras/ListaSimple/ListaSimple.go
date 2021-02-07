package ListaSimple

import "fmt"

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

func (elist *list)GetSize() int{
	auxiliar := elist.first
	var counter int = 0
	for auxiliar != nil {
		counter++
		auxiliar = auxiliar.next
	}
	return counter
}