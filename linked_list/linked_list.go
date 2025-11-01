package my_structs

import (
	"errors"

	node "github.com/kolaabaka/my_structs/node"
)

type linkedList[T any] struct {
	firstNode *node.Node[T]
	lastNode  *node.Node[T]
}

func CreateLinkedList[T any](firstNode *node.Node[T], lastNode *node.Node[T]) *linkedList[T] {
	if firstNode.GetNext() != lastNode {
		firstNode.SetNext(lastNode)
	}

	if lastNode.GetPrevious() != firstNode {
		lastNode.SetPrevious(firstNode)
	}

	return &linkedList[T]{
		firstNode: firstNode,
		lastNode:  lastNode,
	}
}

func (l *linkedList[T]) GetFirst() *node.Node[T] {
	return l.firstNode
}

func (l *linkedList[T]) GetLast() *node.Node[T] {
	return l.lastNode
}

// todo: check which half of the indices contains the input parameter and make two algorithms from beginnig or from ending
func (l *linkedList[T]) Get(index int) (*node.Node[T], error) {
	counter := 0
	node := l.firstNode
	for counter < index {
		if node.GetNext() == nil {
			return nil, errors.New("index out of bound")
		}
		counter++
		node = node.GetNext()
	}
	return node, nil
}

func (l *linkedList[T]) PushLast(node *node.Node[T]) {
	node.SetPrevious(l.lastNode)
	l.lastNode.SetNext(node)

	l.lastNode = node
}

func (l *linkedList[T]) PushFirst(node *node.Node[T]) {
	node.SetNext(l.firstNode)
	l.firstNode.SetPrevious(node)

	l.firstNode = node
}

func (l *linkedList[T]) LinkedListLength() int {
	noda := l.firstNode
	var counter int

	for noda != nil {
		counter++
		noda = noda.GetNext()
	}
	return counter
}
