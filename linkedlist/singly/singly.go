package main

import (
	"errors"
	"log"
)

type Node struct {
	Data interface{}
	Next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Add adds a new value to the list.
func (l *SinglyLinkedList) Add(node *Node) *SinglyLinkedList {
	node.Next = l.head
	l.head = node
	l.length++
	if l.length == 1 {
		l.tail = l.head
	}
	return l
}

// Iterate loops through the list.
func (l SinglyLinkedList) Iterate(f func(index int, node *Node)) {
	for i := 0; i < l.length; i++ {
		f(i, l.head)
		l.head = l.head.Next
	}
}

// Size retrieves the size of the list.
func (l SinglyLinkedList) Size() int {
	return l.length
}

// IsEmpty returns true if the list is empty else false.
func (l SinglyLinkedList) IsEmpty() bool {
	return l.length == 0
}

// GetHead returns the head of the list.
func (l *SinglyLinkedList) GetHead() *Node {
	return l.head
}

// GetTail returns the tail of the list.
func (l *SinglyLinkedList) GetTail() *Node {
	return l.tail
}

// Delete deletes a node from the list.
func (l *SinglyLinkedList) Delete(node *Node) error {
	if l.IsEmpty() {
		return errors.New("List is empty !")
	}
	trav1 := l.head
	trav2 := l.head.Next
	if l.head == node {
		l.head = l.head.Next
		l.length--
		return nil
	}
	for i := 0; i < l.length; i++ {
		if trav2 != nil && trav2 == node {
			if trav2 == l.tail {
				l.tail = trav1
			}
			trav1.Next = trav2.Next
			l.length--
			return nil
		}
		trav1 = trav2
		if trav2 != nil && trav2.Next != nil {
			trav2 = trav2.Next
		}
	}
	return errors.New("Node does not exist !")
}

// Replace replaces an existing node in the list.
func (l *SinglyLinkedList) Replace(old *Node, new *Node) error {
	if l.IsEmpty() {
		return errors.New("List is empty !")
	}
	trav1 := l.head
	trav2 := l.head.Next
	currentNode := l.head
	if trav1 == old {
		new.Next = trav1.Next
		l.head = new
		return nil
	}
	for i := 0; i < l.length; i++ {
		if trav2 == old {
			new.Next = currentNode.Next
			trav1.Next = new
			return nil
		}
		trav1 = trav2
		if trav2.Next != nil {
			trav2 = trav2.Next
		}
	}
	return errors.New("'Old' node does not exist !")
}

// SAMPLE USE.
func main() {
	linkedList := new(SinglyLinkedList)
	node1 := &Node{Data: "hello"}
	node2 := &Node{Data: "world"}
	node3 := &Node{Data: "i am"}
	node4 := &Node{Data: "wisdommatt"}
	node5 := &Node{Data: "from"}
	node6 := &Node{Data: "Nigeria."}
	linkedList.Add(node1).Add(node2).Add(node3).Add(node4).Add(node5).Add(node6)

	log.Println(linkedList.Delete(node1))
	log.Println(linkedList.Replace(node4, &Node{Data: "Am a Gopher"}))

	linkedList.Iterate(func(i int, node *Node) {
		log.Println("current: ", i, node.Data)
	})
}
