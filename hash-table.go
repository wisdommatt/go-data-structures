package datastructures

import (
	"fmt"
	"math"
	"reflect"
)

// HashTable represents a hash table data structure.
//
// this hash table data structure implementation uses
// separate chaining technique to handle collisions.
type HashTable struct {
	size          int
	elementsCount int
	table         []*DoublyLinkedList
}

// HashTableEntry is the object that represents a hash table entry.
type HashTableEntry struct {
	Key   interface{}
	Value interface{}
}

// NewHashTable returns a new hash table data structure.
func NewHashTable(size int) *HashTable {
	ht := &HashTable{
		size:          size,
		elementsCount: 0,
		table:         make([]*DoublyLinkedList, size),
	}
	for i := 0; i < size; i++ {
		ht.table[i] = nil
	}
	return ht
}

// Set sets a new <Key, Value> item in the hash table.
func (h *HashTable) Set(key interface{}, value interface{}) error {
	keyHash, err := h.hash(key)
	if err != nil {
		return err
	}
	keyList := h.table[keyHash]
	if keyList == nil {
		linkedList := NewDoublyLinkedList()
		linkedList.Add(&DoublyLinkedListNode{
			Data: HashTableEntry{Key: key, Value: value},
		})
		h.table[keyHash] = linkedList
		h.elementsCount++
		return nil
	}
	// checking if this key already exist in the hash table.
	//
	// if it already exist, the hash table should update the value
	// in the data table entry.
	var exist bool
	h.table[keyHash].Iterate(func(_ int, node *DoublyLinkedListNode) {
		if node.Data.(HashTableEntry).Key == key {
			exist = true
			node.Data = HashTableEntry{Key: key, Value: value}
		}
	})
	if exist {
		return nil
	}

	h.table[keyHash].Add(&DoublyLinkedListNode{
		Data: HashTableEntry{Key: key, Value: value},
	})
	h.elementsCount++
	return nil
}

// Get retrieves an item from the hash table using the key.
func (h *HashTable) Get(key interface{}) (interface{}, error) {
	keyHash, err := h.hash(key)
	if err != nil {
		return nil, err
	}
	linkedList := h.table[keyHash]
	if linkedList == nil {
		return nil, fmt.Errorf("no item found for key: %d", key)
	}
	var value interface{}
	// this variable is used to check if the key actually exist because a
	// key can have nil as the value (in this case it will be confusing to test
	// if the value exist).
	var exist bool
	linkedList.Iterate(func(_ int, node *DoublyLinkedListNode) {
		entry := node.Data.(HashTableEntry)
		if entry.Key == key {
			value = entry.Value
			exist = true
		}
	})
	if exist {
		return value, nil
	}
	return nil, fmt.Errorf("no item found for key: %d", key)
}

// Delete removes an item from the hash table in key position.
//
// if there is no item at key position, delete does nothing.
func (h *HashTable) Delete(key interface{}) {
	keyHash, err := h.hash(key)
	if err != nil {
		return
	}
	linkedList := h.table[keyHash]
	if linkedList == nil {
		return
	}
	newLinkedListNodes := []*DoublyLinkedListNode{}
	linkedList.Iterate(func(_ int, node *DoublyLinkedListNode) {
		if node.Data.(HashTableEntry).Key == key {
			h.elementsCount--
			return
		}
		newLinkedListNodes = append(newLinkedListNodes, node)
	})
	linkedList.Clear()
	for _, node := range newLinkedListNodes {
		linkedList.Add(node)
	}
}

// Size returns the size of the hash table.
func (h *HashTable) Size() int {
	return h.size
}

// Elements returns the number of elements in the hash table.
func (h *HashTable) Elements() int {
	return h.elementsCount
}

// Iterate iterates through the hash table and executes the callback function
// f for each iteration.
func (h *HashTable) Iterate(f func(key, value interface{})) {
	for _, linkedList := range h.table {
		if linkedList == nil {
			continue
		}
		linkedList.Iterate(func(index int, node *DoublyLinkedListNode) {
			entry := node.Data.(HashTableEntry)
			f(entry.Key, entry.Value)
		})
	}
}

// hash is the hash function for hashing keys.
func (h *HashTable) hash(key interface{}) (int, error) {
	switch key := key.(type) {
	case string:
		return h.hashString(key), nil

	case float64:
		return h.hashNumber(int(key)), nil

	case float32:
		return h.hashNumber(int(key)), nil

	case int64:
		return h.hashNumber(int(key)), nil

	case int32:
		return h.hashNumber(int(key)), nil

	case int16:
		return h.hashNumber(int(key)), nil

	case int:
		return h.hashNumber(key), nil

	default:
		return 0, fmt.Errorf(
			"type of: (%v) cannot be used as key, only string, int and floats are accepted",
			reflect.TypeOf(key),
		)
	}
}

// hashString is a helper function used to hash string key.
func (h *HashTable) hashString(str string) int {
	hash := 0
	for k, s := range str {
		asciiValue := int(s)
		if k == 0 {
			hash += asciiValue + (13 ^ k)
			continue
		}
		hash += asciiValue * (13 ^ k)
	}
	return hash % h.size
}

// hashNumber is a helper function used to hash a number key.
func (h *HashTable) hashNumber(k int) int {
	hash := k * 2654435761 % int(math.Pow(2, 32))
	return hash % h.size
}
