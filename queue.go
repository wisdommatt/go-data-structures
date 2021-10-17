package datastructures

import "errors"

type Queue struct {
	linkedList *DoublyLinkedList
}

// NewQueue returns a new queue data structure.
func NewQueue() *Queue {
	return &Queue{
		linkedList: NewDoublyLinkedList(),
	}
}

// Enqueue adds an item to the tail of the queue.
func (q *Queue) Enqueue(item interface{}) *Queue {
	q.linkedList.AddTail(&DoublyLinkedListNode{Data: item})
	return q
}

// Size returns the size of the queue.
func (q *Queue) Size() int {
	return q.linkedList.Size()
}

// IsEmpty returns true if the queue is empty; else false.
func (q *Queue) IsEmpty() bool {
	return q.linkedList.IsEmpty()
}

// Dequeue removes the first element from the head of the queue.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	head := q.linkedList.GetHead()
	q.linkedList.RemoveHead()
	return head.Data, nil
}

// Peek returns the value of the first element in the queue without
// removing it.
func (q *Queue) Peek() (interface{}, error) {
	headNode := q.linkedList.GetHead()
	if headNode == nil {
		return nil, errors.New("queue is empty")
	}
	return headNode.Data, nil
}

// Iterate iterates through the queue and executes the callback function
// f for each iteration.
func (q *Queue) Iterate(f func(index int, item interface{})) {
	q.linkedList.Iterate(func(index int, node *DoublyLinkedListNode) {
		f(index, node.Data)
	})
}

// Contains returns true if the item is in the queue; else false.
func (q *Queue) Contains(item interface{}) bool {
	var exist bool
	q.Iterate(func(index int, elem interface{}) {
		if item == elem {
			exist = true
			return
		}
	})
	return exist
}
