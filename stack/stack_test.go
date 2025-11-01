package my_structs

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Create[int](0)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Show()

	if stack.IsEmpty() {
		t.Error("push doesn`t work")
	}

	index, err := stack.Search(1)
	if err != nil {
		t.Error("something go wrong in Search() function")
	}
	if index != 2 {
		t.Errorf("Search() doesn`t work, index value = %d", index)
	}

	_, err = stack.Pop(10000)
	if err == nil {
		t.Error("Pop(10000) must return error")
	}

	testSliceOne, err := stack.Pop(3)
	if err != nil {
		t.Error("something go wrong in Pop() function")
	}
	if len(testSliceOne) != 3 {
		t.Errorf("Pop() didn`t rerurn reqired quantity, it returned %d", len(testSliceOne))
	}

	testSliceTwo, err := stack.Peek(3)
	if err != nil {
		t.Error("something go wrong in Pop() function")
	}
	if len(testSliceTwo) != 3 {
		t.Errorf("Peek() didn`t rerurn reqired quantity, it returned %d", len(testSliceTwo))
	}
	if !stack.IsEmpty() {
		t.Error("Peek() didn`t remove elements from stack")
	}
}
