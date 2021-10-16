package datastructures

import (
	"testing"
)

func TestDoublyLinkedList_Add(t *testing.T) {
	list := NewDoublyLinkedList()
	node := &DoublyLinkedListNode{Data: "Hello World"}
	list.Add(node)
	if list.GetHead() != node {
		t.Errorf("GetHead(): expected = %v, got = %v", node, list.GetHead())
	}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.Add(node2)
	if list.GetTail() != node2 {
		t.Errorf("GetTail(): expected = %v, got = %v", node2, list.GetTail())
	}
	node3 := &DoublyLinkedListNode{Data: "Third Node"}
	list.Add(node3)
	if list.GetTail() != node3 {
		t.Errorf("GetTail(): expected = %v, got = %v", node3, list.GetTail())
	}
}

func TestDoublyLinkedList_Clear(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Size(): expected = %v, got = %v", 0, list.Size())
	}
	node := &DoublyLinkedListNode{Data: "First Node"}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.Add(node).Add(node2)
	if list.Size() != 2 {
		t.Errorf("Size(): expected = %v, got = %v", 2, list.Size())
	}
	list.Clear()
	if list.Size() != 0 {
		t.Errorf("Size(): expected = %v, got = %v", 0, list.Size())
	}
	// readding the nodes
	list.Add(node).Add(node2)
	if list.Size() != 2 {
		t.Errorf("Size(): expected = %v, got = %v", 2, list.Size())
	}
}

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	list := NewDoublyLinkedList()
	node := &DoublyLinkedListNode{Data: "First Node"}
	list.AddHead(node)
	if list.GetHead() != node {
		t.Errorf("GetHead(): expected = %v, got = %v", node, list.GetHead())
	}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.AddHead(node2)
	if list.GetHead() != node2 {
		t.Errorf("GetHead(): expected = %v, got = %v", node2, list.GetHead())
	}
	if list.GetTail() != node {
		t.Errorf("GetTail(): expected = %v, got = %v", node, list.GetTail())
	}
}

func TestDoublyLinkedList_AddTail(t *testing.T) {
	list := NewDoublyLinkedList()
	node := &DoublyLinkedListNode{Data: "First Node"}
	list.AddTail(node)
	if list.GetTail() != node {
		t.Errorf("GetTail(): expected = %v, got = %v", node, list.GetTail())
	}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.AddTail(node2)
	if list.GetTail() != node2 {
		t.Errorf("GetTail(): expected = %v, got = %v", node2, list.GetTail())
	}
	if list.GetHead() != node {
		t.Errorf("GetHead(): expected = %v, got = %v", node, list.GetHead())
	}
}

func TestDoublyLinkedList_Iterate(t *testing.T) {
	tests := []struct {
		name  string
		nodes []*DoublyLinkedListNode
	}{
		{
			name: "7 nodes",
			nodes: []*DoublyLinkedListNode{
				{Data: "1"},
				{Data: "2"},
				{Data: "3"},
				{Data: "4"},
				{Data: "5"},
				{Data: "6"},
				{Data: "7"},
			},
		},
		{
			name: "10 nodes",
			nodes: []*DoublyLinkedListNode{
				{Data: "1"},
				{Data: "2"},
				{Data: "3"},
				{Data: "4"},
				{Data: "5"},
				{Data: "6"},
				{Data: "7"},
				{Data: "8"},
				{Data: "9"},
				{Data: "10"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewDoublyLinkedList()
			for _, n := range tt.nodes {
				l.Add(n)
			}
			l.Iterate(func(index int, node *DoublyLinkedListNode) {
				if tt.nodes[index] != node {
					t.Errorf("index: %d, expected = %v, got = %v", index, tt.nodes[index], node)
				}
			})
		})
	}
}

func TestDoublyLinkedList_RemoveHead(t *testing.T) {
	list := NewDoublyLinkedList()
	list.RemoveHead()
	if list.GetHead() != nil {
		t.Errorf("GetHead(): expected = %v, got = %v", nil, list.GetHead())
	}
	node := &DoublyLinkedListNode{Data: "First Node"}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.AddHead(node)
	list.RemoveHead()
	if list.GetHead() != nil {
		t.Errorf("GetHead(): expected = %v, got = %v", nil, list.GetHead())
	}
	list.AddHead(node).AddHead(node2)
	if list.GetHead() != node2 {
		t.Errorf("GetHead(): expected = %v, got = %v", node2, list.GetHead())
	}
	list.RemoveHead()
	if list.GetHead() != node {
		t.Errorf("GetHead(): expected = %v, got = %v", node, list.GetHead())
	}
}

func TestDoublyLinkedList_RemoveTail(t *testing.T) {
	list := NewDoublyLinkedList()
	list.RemoveTail()
	if list.GetTail() != nil {
		t.Errorf("GetTail(): expected = %v, got = %v", nil, list.GetTail())
	}
	node := &DoublyLinkedListNode{Data: "First Node"}
	node2 := &DoublyLinkedListNode{Data: "Second Node"}
	list.AddTail(node)
	list.RemoveTail()
	if list.GetTail() != nil {
		t.Errorf("GetTail(): expected = %v, got = %v", nil, list.GetTail())
	}
	list.AddTail(node).AddTail(node2)
	if list.GetTail() != node2 {
		t.Errorf("GetTail(): expected = %v, got = %v", node2, list.GetTail())
	}
	list.RemoveTail()
	if list.GetTail() != node {
		t.Errorf("GetTail(): expected = %v, got = %v", node, list.GetTail())
	}
}
