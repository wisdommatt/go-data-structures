package datastructures

import (
	"errors"
)

type Stack struct {
	linkedList *DoublyLinkedList
	length     int
}

// NewStack returns a new stack data structure.
func NewStack() *Stack {
	return &Stack{
		linkedList: NewDoublyLinkedList(),
	}
}

// Size returns the current size of the stack.
func (s *Stack) Size() int {
	return s.linkedList.Size()
}

// IsEmpty returns true if the stack is empty else false.
func (s *Stack) IsEmpty() bool {
	return s.linkedList.Size() == 0
}

// Push adds a new item to the stack.
func (s *Stack) Push(item interface{}) *Stack {
	s.linkedList.AddHead(&DoublyLinkedListNode{Data: item})
	s.length++
	return s
}

// Pop removes the top element from the stack.
func (s *Stack) Pop() (interface{}, error) {
	headNode := s.linkedList.GetHead()
	if headNode == nil {
		return nil, errors.New("stack is empty")
	}
	s.linkedList.RemoveHead()
	return headNode.Data, nil
}

// Peek returns the top element in the stack without removing it.
func (s *Stack) Peek() (interface{}, error) {
	headNode := s.linkedList.GetHead()
	if headNode == nil {
		return nil, errors.New("stack is empty")
	}
	return headNode.Data, nil
}

// Iterate iterates through stack and executes the callback function
// f for each iteration.
func (s *Stack) Iterate(f func(index int, item interface{})) {
	s.linkedList.Iterate(func(index int, node *DoublyLinkedListNode) {
		f(index, node.Data)
	})
}

// Contains returns true if the item is in the stack; else false.
func (s *Stack) Contains(item interface{}) bool {
	var exist bool
	s.Iterate(func(index int, elem interface{}) {
		if item == elem {
			exist = true
			return
		}
	})
	return exist
}
