package main

import (
	"errors"
	"log"
)

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *DoublyLinkedList) Add(node *Node) *DoublyLinkedList {
	node.Previous = l.head
	if l.head != nil {
		l.head.Next = node
	}
	l.head = node
	l.length++
	if l.length == 1 {
		l.tail = l.head
	}
	return l
}

func (l DoublyLinkedList) Iterate(f func(index int, node *Node)) {
	for i := 0; i < l.length; i++ {
		f(i, l.head)
		l.head = l.head.Previous
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
func (l *DoublyLinkedList) GetHead() *Node {
	return l.head
}

// GetTail returns the tail of the list.
func (l *DoublyLinkedList) GetTail() *Node {
	return l.tail
}

// Delete deletes a node from the list.
func (l *DoublyLinkedList) Delete(node *Node) error {
	if l.IsEmpty() {
		return errors.New("List is empty !")
	}
	if l.head == node {
		if l.head.Previous != nil {
			l.head.Previous.Next = l.head.Next
		}
		l.head = l.head.Previous
		l.length--
		return nil
	}
	current := l.head
	for i := 0; i < l.length; i++ {
		if current == node {
			current.Previous.Next = current.Next
			if current.Next != nil {
				current.Next.Previous = current.Previous
			}
			l.length--
			return nil
		}
		current = current.Previous
	}
	return errors.New("Node does not exist !")
}

// Replace replaces an existing node in the list.
func (l *DoublyLinkedList) Replace(old *Node, new *Node) error {
	if l.IsEmpty() {
		return errors.New("List is empty !")
	}
	if l.head == old {
		if l.head.Previous != nil {
			l.head.Previous.Next = new
			new.Previous = l.head.Previous
		}
		if l.head.Next != nil {
			l.head.Next.Previous = new
			new.Next = l.head.Next
		}
		return nil
	}
	current := l.head
	for i := 0; i < l.length; i++ {
		if current == old {
			current.Previous.Next = new
			if current.Next != nil {
				current.Next.Previous = new
			}
			new.Previous = current.Previous
			new.Next = current.Next
			return nil
		}
		current = current.Previous
	}
	return errors.New("Node does not exist !")
}

func main() {
	node1 := &Node{Data: "1"}
	node2 := &Node{Data: "2"}
	node3 := &Node{Data: "3"}
	node4 := &Node{Data: "4"}
	node5 := &Node{Data: "5"}
	linkedList := new(DoublyLinkedList)
	linkedList.Add(node1).Add(node2).Add(node3).Add(node4).Add(node5)

	wakandaNode := &Node{Data: "Wakanda"}
	log.Println(linkedList.Replace(node3, wakandaNode))
	log.Println(linkedList.Delete(node5))

	linkedList.Iterate(func(index int, node *Node) {
		log.Println(index, node)
	})
}
