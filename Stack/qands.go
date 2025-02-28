package qands

import (
	"fmt"
	"github.com/dataStructureAndalgorithms/ll"
)

type Stack[T any] struct { 
    ll.LinkedList[T]
}

func (s *Stack[T]) Push(val T) {
	s.InsertAtHead(val)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.Length() == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	top, err := s.Top()
	if err != nil {
		return zero, err
	}
	s.DeleteHead()
	return top, nil
}

func (s *Stack[T]) Top() (T, error) {
	var zero T
	if s.Length() == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	val, err := s.GetVal(0)
	if err != nil {
		return zero, err
	}
	return val,nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Size() int {
    return s.Length()
}

func (s *Stack[T]) Clear() {
    s.Head = nil
}
