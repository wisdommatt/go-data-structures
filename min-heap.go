package datastructures

import (
	"errors"
	"fmt"
)

// MinHeap represents the min-heap data structure.
type MinHeap struct {
	hashTable map[float64][]int
	items     []float64
	length    int
}

// NewMinHeap returns a new min-heap data structure.
func NewMinHeap() *MinHeap {
	return &MinHeap{
		hashTable: make(map[float64][]int),
	}
}

// Insert adds a new item to the heap.
func (h *MinHeap) Insert(item float64) *MinHeap {
	h.items = append(h.items, item)
	h.length++
	if _, ok := h.hashTable[item]; !ok {
		h.hashTable[item] = []int{}
	}
	h.hashTable[item] = append(h.hashTable[item], h.length-1)
	h.bubbleUpFromIndex(h.length - 1)
	return h
}

func (h *MinHeap) bubbleUpFromIndex(index int) {
	if h.length < 2 {
		return
	}
	parentIndex := h.getNodeParentIndex(index)
	if h.items[index] < h.items[parentIndex] {
		h.swap(parentIndex, index, true)
	}
}

func (h *MinHeap) getNodeParentIndex(nodeIndex int) int {
	// indexes 0 - 3 are edge cases.
	if nodeIndex < 3 {
		return 0
	}
	if nodeIndex == 3 {
		return 1
	}
	parentIndex := ((nodeIndex - 2) / 2)
	if (nodeIndex % 2) != 0 {
		parentIndex = ((nodeIndex - 1) / 2)
	}
	return parentIndex
}

// Poll removes the root element from the heap.
func (h *MinHeap) Poll() (float64, error) {
	return h.removeAtPosition(0)
}

// removeAtPosition removes an element from the specified position.
func (h *MinHeap) removeAtPosition(position int) (float64, error) {
	if h.length == 0 {
		return 0, errors.New("heap is empty")
	}
	itemAtPos := h.items[position]
	h.removePositionFromHashTable(itemAtPos, position)
	if h.length == 1 {
		h.items = nil
		h.length--
		return itemAtPos, nil
	}
	h.items[position] = h.items[h.length-1]
	h.items = h.items[:h.length-1]
	h.length--
	// checking if the element is the last element in the heap.
	//
	// if the element is the last element in the heap, there is no
	// need to bubble down.
	if position != h.length {
		h.removePositionFromHashTable(h.items[position], h.length)
		h.hashTable[h.items[position]] = append(h.hashTable[h.items[position]], position)
		h.bubbleDownFromIndex(position)
	}
	return itemAtPos, nil
}

func (h *MinHeap) bubbleDownFromIndex(index int) {
	if h.length < 2 {
		return
	}
	lcIndex, rcIndex := h.getNodeChildrenIndexes(index)
	if h.items[lcIndex] < h.items[index] && h.items[lcIndex] < h.items[rcIndex] {
		h.swap(index, lcIndex, false)
		return
	}
	if h.items[rcIndex] < h.items[index] {
		h.swap(index, rcIndex, false)
	}
}

func (h *MinHeap) removePositionFromHashTable(item float64, position int) {
	arr := h.hashTable[item]
	newArr := []int{}
	for _, v := range arr {
		if v != position {
			newArr = append(newArr, v)
		}
	}
	h.hashTable[item] = newArr
}

func (h *MinHeap) getNodeChildrenIndexes(nodeIndex int) (int, int) {
	leftChildIndex := (2 * nodeIndex) + 1
	rightChildIndex := (2 * nodeIndex) + 2
	if (h.length - 1) < leftChildIndex {
		return nodeIndex, nodeIndex
	}
	if (h.length - 1) < rightChildIndex {
		return leftChildIndex, nodeIndex
	}
	return leftChildIndex, rightChildIndex
}

func (h *MinHeap) swap(index1, index2 int, bubbleUp bool) {
	parentValue := h.items[index1]
	childValue := h.items[index2]
	h.items[index1] = childValue
	h.items[index2] = parentValue
	// updating the indexes in the hash table.
	h.removePositionFromHashTable(parentValue, index1)
	h.removePositionFromHashTable(childValue, index2)
	h.hashTable[parentValue] = append(h.hashTable[parentValue], index2)
	h.hashTable[childValue] = append(h.hashTable[childValue], index1)

	if bubbleUp {
		newParentIndex := h.getNodeParentIndex(index1)
		if h.items[index1] < h.items[newParentIndex] {
			h.swap(newParentIndex, index1, true)
		}
	}

	bubbleDown := !bubbleUp
	if bubbleDown {
		lcIndex, rcIndex := h.getNodeChildrenIndexes(index2)
		if h.items[lcIndex] < h.items[index2] && h.items[lcIndex] < h.items[rcIndex] {
			h.swap(index2, lcIndex, false)
		}
		if h.items[rcIndex] < h.items[index2] {
			h.swap(index2, rcIndex, false)
		}
	}
}

// Remove removes an item from the heap.
func (h *MinHeap) Remove(item float64) (float64, error) {
	positions, ok := h.hashTable[item]
	if !ok || len(positions) == 0 {
		return 0, fmt.Errorf("%f is not in heap", item)
	}
	// remove the last position from the hast table.
	lastPosition := positions[len(positions)-1]
	return h.removeAtPosition(lastPosition)
}

// Contains returns true if the item is in the heap, else false.
func (h *MinHeap) Contains(item float64) bool {
	positions, ok := h.hashTable[item]
	if !ok || len(positions) == 0 {
		return false
	}
	return true
}

// Size returns the size of the heap.
func (h *MinHeap) Size() int {
	return h.length
}

// Peek returns the smallest item of the heap without removing it.
func (h *MinHeap) Peek() (float64, error) {
	if h.Size() == 0 {
		return 0, errors.New("heap is empty")
	}
	return h.items[0], nil
}

// GetList returns the heap items as a list.
//
// time complexity: 0(1)
func (h *MinHeap) GetList() []float64 {
	return h.items
}
