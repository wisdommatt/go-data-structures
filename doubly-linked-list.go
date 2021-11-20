package datastructures

type DoublyLinkedListNode struct {
	Data     interface{}
	Next     *DoublyLinkedListNode
	Previous *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head   *DoublyLinkedListNode
	tail   *DoublyLinkedListNode
	length int
}

// NewDoublyLinkedList returns a new doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Add adds a new node to the tail doubly linked list.
func (l *DoublyLinkedList) Add(node *DoublyLinkedListNode) *DoublyLinkedList {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
		l.length++
		return l
	}
	node.Previous = l.tail
	l.tail.Next = node
	l.tail = node
	l.length++
	return l
}

// AddHead adds a new node to the head of the doubly linked list.
func (l *DoublyLinkedList) AddHead(node *DoublyLinkedListNode) *DoublyLinkedList {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
		l.length++
		return l
	}
	node.Next = l.head
	l.head.Previous = node
	l.head = node
	l.length++
	return l
}

// AddTail adds a new node to the tail of the linked list.
func (l *DoublyLinkedList) AddTail(node *DoublyLinkedListNode) *DoublyLinkedList {
	if l.IsEmpty() {
		l.head = node
		l.tail = node
		l.length++
		return l
	}
	node.Previous = l.tail
	l.tail.Next = node
	l.tail = node
	l.length++
	return l
}

// Clear clears all the values from the linked list.
func (l *DoublyLinkedList) Clear() {
	if l.IsEmpty() {
		return
	}
	trav := l.head
	for trav != nil {
		next := trav.Next
		trav.Next = nil
		trav.Previous = nil
		trav = next
	}
	l.length = 0
}

// Iterate iterates through the doubly linked list and executes the callback function
// f for each iteration.
//
// TODO(wisdommatt): include iterator with Stop method in the f callback function
// parameters. The reason for this is to stop the iteration if the remaining items
// are not needed.
func (l *DoublyLinkedList) Iterate(f func(index int, node *DoublyLinkedListNode)) {
	trav := l.head
	index := 0
	for trav != nil {
		next := trav.Next
		f(index, trav)
		trav = next
		index++
	}
}

// Size retrieves the size of the list.
func (l DoublyLinkedList) Size() int {
	return l.length
}

// IsEmpty returns true if the list is empty else false.
func (l DoublyLinkedList) IsEmpty() bool {
	return l.length == 0
}

// GetHead returns the head of the list.
func (l *DoublyLinkedList) GetHead() *DoublyLinkedListNode {
	return l.head
}

// GetTail returns the tail of the list.
func (l *DoublyLinkedList) GetTail() *DoublyLinkedListNode {
	return l.tail
}

// RemoveHead removes a node from the head of the doubly linked list.
func (l *DoublyLinkedList) RemoveHead() *DoublyLinkedList {
	if l.IsEmpty() {
		return l
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return l
	}
	l.head = l.head.Next
	l.length--
	return l
}

// RemoveTail removes a node from the tail of the doubly linked list.
func (l *DoublyLinkedList) RemoveTail() *DoublyLinkedList {
	if l.IsEmpty() {
		return l
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return l
	}
	l.tail = l.tail.Previous
	l.length--
	return l
}
