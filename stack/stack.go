package my_structs

import (
	"errors"
	"fmt"
	"sync"
)

type LIFO[T comparable] interface {
	Push(element T) bool
	Pop(amount int) ([]T, error)
	Peek(amount int) ([]T, error)
	IsEmpty() bool
	Search(element T) (int, error)
	Show()
}

type stack[T comparable] struct {
	stack  []T
	locker sync.Mutex
}

func Create[T comparable](length int) *stack[T] {
	return &stack[T]{
		stack:  make([]T, length),
		locker: sync.Mutex{},
	}
}

func (s *stack[T]) Push(element T) bool {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.stack = append(s.stack, element)
	return true
}

func (s *stack[T]) Pop(amount int) ([]T, error) {
	if amount > len(s.stack) {
		return nil, errors.New("amount more than stack size")
	}

	s.locker.Lock()
	defer s.locker.Unlock()

	values := make([]T, 0)

	values = append(values, s.stack[len(s.stack)-amount:]...)

	return values, nil
}

func (s *stack[T]) Peek(amount int) ([]T, error) {
	if amount > len(s.stack) {
		return nil, errors.New("amount more than stack size")
	}

	s.locker.Lock()
	defer s.locker.Unlock()

	values := make([]T, 0)

	values = append(values, s.stack[len(s.stack)-amount:]...)

	s.stack = s.stack[:len(s.stack)-amount]

	return values, nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *stack[T]) Search(element T) (int, error) {
	for i := len(s.stack) - 1; i >= 0; i-- {
		if s.stack[i] == element {
			return len(s.stack) - i - 1, nil
		}
	}
	return -1, errors.New("no such element in stack")
}

func (s *stack[T]) Show() {
	fmt.Print(s.stack)
}
