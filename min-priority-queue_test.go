package datastructures

import (
	"testing"
)

func TestMinPriorityQueue_Enqueue(t *testing.T) {
	type args struct {
		item float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "inserting item to priority queue",
			args: args{item: 109},
		},
		{
			name: "inserting another item to the priority queue",
			args: args{item: 10989},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewMinPriorityQueue()
			q.Enqueue(tt.args.item)
			if q.Size() != 1 {
				t.Errorf("MinPriorityQueue.Enqueue() = %v, want %v", q.Size(), 1)
			}
		})
	}
}

func TestMinPriorityQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		items   []float64
		want    float64
		wantErr bool
	}{
		{
			name:  "dequeuing an item from 6 items queue",
			items: []float64{6, 7, 3, 6, 2, 1},
			want:  1,
		},
		{
			name:  "dequeuing an item from 1 item queue",
			items: []float64{300},
			want:  300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewMinPriorityQueue()
			for _, i := range tt.items {
				q.Enqueue(i)
			}
			got, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("MinPriorityQueue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinPriorityQueue.Dequeue() = %v, want %v", got, tt.want)
			}
			if q.Size() != len(tt.items)-1 {
				t.Errorf("MinPriorityQueue.Dequeue() = %v, want %v", q.Size(), len(tt.items)-1)
			}
		})
	}
}

func TestMinPriorityQueue_Contains(t *testing.T) {
	type args struct {
		item float64
	}
	tests := []struct {
		name  string
		items []float64
		args  args
		want  bool
	}{
		{
			name:  "invalid item",
			items: []float64{440, 2009444, 990029},
			args:  args{item: 445},
			want:  false,
		},
		{
			name:  "valid item",
			items: []float64{111, 222, 333, 444},
			args:  args{item: 333},
			want:  true,
		},
		{
			name: "invalid item in empty queue",
			args: args{item: 1233},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewMinPriorityQueue()
			for _, i := range tt.items {
				q.Enqueue(i)
			}
			if got := q.Contains(tt.args.item); got != tt.want {
				t.Errorf("MinPriorityQueue.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPriorityQueue_Iterate(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "queue with 5 items",
			items: []float64{8, 4, 5, 3, 6},
			want:  []float64{3, 4, 5, 8, 6},
		},
		{
			name: "empty queue",
		},
		{
			name:  "queue with 1 item",
			items: []float64{7},
			want:  []float64{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewMinPriorityQueue()
			for _, i := range tt.items {
				q.Enqueue(i)
			}
			q.Iterate(func(index int, item float64) {
				if tt.want[index] != item {
					t.Errorf("index: %d, expected = %v, got = %v", index, tt.want[index], item)
				}
			})
		})
	}
}
