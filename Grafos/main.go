package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Vertice struct {
	Numero     int
	Adyacentes *list.List
}
type ListaAdyacencia struct {
	Lista *list.List
}

func NewVertice(num int) *Vertice {
	a := list.New()
	return &Vertice{num, a}
}
func NewListaAdyacencia() *ListaAdyacencia {
	a := list.New()
	return &ListaAdyacencia{a}
}
func (this *ListaAdyacencia) getVertice(i int) *Vertice {
	for e := this.Lista.Front(); e != nil; e = e.Next() {
		a := e.Value
		b := a.(*Vertice)
		if b.Numero == i {
			return b
		}
	}
	return nil
}
func (this *ListaAdyacencia) Insertar(i int) {
	if this.getVertice(i) == nil {
		n := NewVertice(i)
		this.Lista.PushBack(n)
	} else {
		fmt.Println("Elemento ya agregado")
	}
}
func (this *ListaAdyacencia) enlazar(a int, b int) {
	var origen *Vertice
	var destino *Vertice
	origen = this.getVertice(a)
	destino = this.getVertice(b)
	if origen == nil || destino == nil {
		fmt.Println("No se encontro el vertice")
		return
	}
	origen.Adyacentes.PushFront(destino)
	destino.Adyacentes.PushFront(origen)
}
func (this *Vertice) Contiene(num int) bool {
	for e := this.Adyacentes.Front(); e != nil; e = e.Next() {
		if e.Value == num {
			return true
		}
	}
	return false
}
func (this *ListaAdyacencia) dibujar() {
	aux := list.New()
	var sc strings.Builder
	fmt.Fprintf(&sc, "digraph G{\n")
	for e := this.Lista.Front(); e != nil; e = e.Next() {
		tmp := (e.Value).(*Vertice)
		if contiene(aux, tmp) == false {
			aux.PushBack(tmp)
			fmt.Fprintf(&sc, "node%p[label=\"%v\"]\n", &(*tmp), tmp.Numero)
		}
		for temporal := tmp.Adyacentes.Front(); temporal != nil; temporal = temporal.Next() {
			verticetemporal := (temporal.Value).(*Vertice)
			fmt.Fprintf(&sc, "node%p->node%p\n", &(*tmp), &(*verticetemporal))
			if contiene(aux, verticetemporal) {

			} else {
				aux.PushBack(verticetemporal)
				fmt.Fprintf(&sc, "node%p[label=\"%v\"]\n", &(*verticetemporal), verticetemporal.Numero)

			}
		}
	}
	fmt.Fprintf(&sc, "}")
	fmt.Println(sc.String())
}
func contiene(buscando *list.List, elemento *Vertice) bool {
	for e := buscando.Front(); e != nil; e = e.Next() {
		if (e.Value).(*Vertice) == elemento {
			return true
		}
	}
	return false
}
func main() {
	nueva := NewListaAdyacencia()
	nueva.Insertar(1)
	nueva.Insertar(2)
	nueva.Insertar(3)
	nueva.Insertar(4)
	nueva.Insertar(5)

	nueva.enlazar(1, 2)
	nueva.enlazar(1, 3)
	nueva.enlazar(3, 2)
	nueva.enlazar(2, 4)
	nueva.enlazar(4, 5)
	nueva.enlazar(3, 5)
	nueva.dibujar()
}
