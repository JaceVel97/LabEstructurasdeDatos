package Pila

type node_stack struct {
	next *node_stack
	value int
}

type stack struct {
	first *node_stack
}

func NewStack() *stack {
	return &stack{nil}
}


func (estack *stack) Push(value int) {
	newNode := &node_stack{nil, -1}
	newNode.value = value
	if estack.first == nil {
		estack.first = newNode
	} else {
		newNode.next = estack.first
		estack.first = newNode
	}
}

func (estack *stack) Peek() int{
	return estack.first.value
}

func (estack *stack) Pop() int {
	var value_stack int = estack.first.value

	if estack.first.next == nil {
		estack.first = nil
	} else {
		estack.first = estack.first.next
	}

	return value_stack
}

func (estack *stack)GetSize() int{
	auxiliar := estack.first
	var counter int = 0
	for auxiliar != nil {
		counter++
		auxiliar = auxiliar.next
	}
	return counter
}