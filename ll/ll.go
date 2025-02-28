package ll

import (
	"fmt"
)

type LinkedList[T any] struct {
	Head *node[T]
}

type node[T any] struct {
	data T
	next *node[T]
}

func (l *LinkedList[T]) Insert(data T) *node[T] {
	newNode := &node[T]{data: data}
	if l.Head == nil {
		l.Head = newNode
		return newNode
	}
	current := l.Head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return newNode
}

func (l *LinkedList[T]) InsertAtHead(data T) {
	newNode := &node[T]{data: data}
	newNode.next = l.Head
	l.Head = newNode
}
func (l *LinkedList[T]) InsertAtTail(data T) {
	newNode := &node[T]{data: data}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *LinkedList[T]) InsertAtIndex(data T, i int) error {
	if i < 1 {
		return fmt.Errorf("use InsertAtHead if you want to insert at the start of the list")
	}
	if i >= l.Length() {
		return fmt.Errorf("index %d out of range", i)
	}
	newNode := &node[T]{data: data}
	nextNode, err := l.getNode(i)
	if err != nil {
		return err
	}
	prevNode, err := l.getNode(i - 1)
	if err != nil {
		return err
	}
	newNode.next = nextNode
	prevNode.next = newNode
	return nil
}

func (l *LinkedList[T]) DeleteTail() error {
	if l.Head == nil {
		return fmt.Errorf("the list is already empty")
	}
	if l.Head.next == nil {
		l.Head = nil
		return nil
	}
	newTail, err := l.getNode(l.Length() - 2)
	if err != nil {
		return err
	}
	newTail.next = nil
	return nil
}

func (l *LinkedList[T]) DeleteAtIndex(index int) error {
	if l.Head == nil {
		return fmt.Errorf("the list is already empty")
	}
	if index < 0 || index >= l.Length() {
		return fmt.Errorf("index %d out of range", index)
	}
	if index == 0 {
		l.Head = l.Head.next
		return nil
	}
	prevNode, err := l.getNode(index - 1)
	if err != nil {
		return err
	}
	if prevNode.next == nil {
		return fmt.Errorf("no node to delete at index %d", index)
	}
	prevNode.next = prevNode.next.next
	return nil
}

func (l *LinkedList[T]) DeleteHead() error {
    if l.Head == nil {
        return fmt.Errorf("list is already empty")
    }
    l.Head = l.Head.next
    return nil
}


func (l *LinkedList[T]) Length() int {
	length := 0
	for current := l.Head; current != nil; current = current.next {
		length++
	}
	return length
}

func (l *LinkedList[T]) getNode(index int) (*node[T], error) {
	if index >= l.Length() || index < 0 {
		return nil, fmt.Errorf("index out of range: %d", index)
	}
	current := l.Head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current, nil
}

func (l *LinkedList[T]) GetVal(i int) (T, error) {
	n, err := l.getNode(i)
	if err != nil {
		var zero T
		return zero, err
	}
	return n.data, nil
}

func (l *LinkedList[T]) PrintList() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Print("\n")
}
