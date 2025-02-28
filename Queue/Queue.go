package Queue

import (
	"fmt"

	"github.com/dataStructureAndalgorithms/ll"
)

type Queue[T any] struct {
	ll.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(val T) {
	q.InsertAtTail(val)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.Length() == 0 {
		return zero, fmt.Errorf("queue is empty")
	}
	peek := q.Peek()

	q.DeleteHead()
	return peek, nil
}

func (q *Queue[T]) Peek() T {
	var zero T
	if q.Length() == 0 {
		return zero
	}
	val, err := q.GetVal(0)
	if err != nil {
		return zero
	}
	return val
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}
