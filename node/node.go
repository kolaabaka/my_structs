package my_structs

type Node[T any] struct {
	previous  *Node[T]
	next      *Node[T]
	container T
}

func (n *Node[T]) SetValue(element T) {
	n.container = element
}

func (n *Node[T]) GetValue() T {
	return n.container
}

func (n *Node[T]) GetPrevious() *Node[T] {
	return n.previous
}

func (n *Node[T]) GetNext() *Node[T] {
	return n.next
}

func (n *Node[T]) SetPrevious(noda *Node[T]) {
	n.previous = noda
}

func (n *Node[T]) SetNext(noda *Node[T]) {
	n.next = noda
}
