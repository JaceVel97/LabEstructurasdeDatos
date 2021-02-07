package ListaDoblementeEnlazada

import "fmt"

type nodeD struct {
	next *nodeD
	previous *nodeD
	value int
}

type listD struct {
	first *nodeD
	last *nodeD
}

func NewList() *listD {
	return &listD{nil,nil}
}

func (elist_d *listD)Insert(value int) {
	newNode := &nodeD{nil, nil, value}
	if elist_d.first == nil {
		elist_d.last = newNode
		elist_d.first = elist_d.last
	} else {
		newNode.previous = elist_d.last
		elist_d.last.next = newNode
		elist_d.last = newNode
	}
}


func (elist_d *listD)ShowData() {
	auxiliar := elist_d.first
	for auxiliar != nil {
		fmt.Println("Valor ", auxiliar.value)
		auxiliar = auxiliar.next
	}
}

func (elist_d *listD)ShowDataBack() {
	auxiliar := elist_d.last
	for auxiliar != nil {
		fmt.Println("Valor ", auxiliar.value)
		auxiliar = auxiliar.previous
	}
}

func (elist_d *listD)GetSize() int{
	auxiliar := elist_d.first
	var counter int = 0
	for auxiliar != nil {
		counter++
		auxiliar = auxiliar.next
	}
	return counter
}