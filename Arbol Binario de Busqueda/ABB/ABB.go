package abb

import "fmt"

type abb struct {
	root *nodeAbb
}

func NewTree() *abb{
	return &abb{nil}
}

func (abb *abb) Insert (value int){
	abb.root = abb.root.InsertData(value)
}

func (tree *nodeAbb) InsertData(value int) *nodeAbb{
	if tree == nil {
		return &nodeAbb{value, nil, nil}
	} else if value < tree.value{
		tree.left = tree.left.InsertData(value)
	} else {
		tree.right = tree.right.InsertData(value)
	}
	return tree
}

func (abb *abb) GetOrder(option int) {
	if option == 0 {
		abb.root.PreOrder()
	} else if option == 1 {
		abb.root.InOrder()
	} else {
		abb.root.PostOrder()
	}
}

func (tree *nodeAbb) PreOrder() {
	if tree != nil {
		fmt.Println(tree.value)
		tree.left.PreOrder()
		tree.right.PreOrder()
	}
}

func (tree *nodeAbb) InOrder() {
	if tree != nil {
		tree.left.InOrder()
		fmt.Println(tree.value)
		tree.right.InOrder()
	}
}

func (tree *nodeAbb) PostOrder() {
	if tree != nil {
		tree.left.PostOrder()
		tree.right.PostOrder()
		fmt.Println(tree.value)
	}
}

func (abb *abb) BuscarBool(value int) bool {
	if abb.root.BuscarRecursivo(value) != nil {
		return true
	} else {
		return false
	}
}

func (abb *abb) BuscarNode(value int) *nodeAbb {
	return abb.root.BuscarRecursivo(value)
}

func (tree *nodeAbb) BuscarRecursivo(value int) *nodeAbb{
	if tree == nil {
		return nil
	} else if value == tree.value {
		return tree
	} else if value < tree.value {
		return tree.left.BuscarRecursivo(value)
	} else {
		return tree.right.BuscarRecursivo(value)
	}
}