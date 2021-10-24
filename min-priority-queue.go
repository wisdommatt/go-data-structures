package datastructures

// MinPriorityQueue represents a min-priority queue data
// structure.
type MinPriorityQueue struct {
	minHeap *MinHeap
}

// NewMinPriorityQueue returns a min-priority queue data structure.
func NewMinPriorityQueue() *MinPriorityQueue {
	return &MinPriorityQueue{
		minHeap: NewMinHeap(),
	}
}

// Enqueue adds an item to the priority queue.
func (q *MinPriorityQueue) Enqueue(item float64) *MinPriorityQueue {
	q.minHeap.Insert(item)
	return q
}

// Dequeue removes the smallest item from the priority queue.
func (q *MinPriorityQueue) Dequeue() (float64, error) {
	return q.minHeap.Poll()
}

// Size returns the size of the priority queue.
func (q *MinPriorityQueue) Size() int {
	return q.minHeap.Size()
}

// Contains returns true if the item is in the queue; else false.
func (q *MinPriorityQueue) Contains(item float64) bool {
	return q.minHeap.Contains(item)
}

// Iterate iterates through the queue and executes the callback function
// f for each iteration.
func (q *MinPriorityQueue) Iterate(f func(index int, item float64)) {
	for k, v := range q.minHeap.GetList() {
		f(k, v)
	}
}
