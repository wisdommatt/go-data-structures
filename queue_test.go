package datastructures

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name  string
		items []interface{}
		want  int
	}{
		{
			name:  "testcase with 3 items",
			items: []interface{}{2, 3, 4},
			want:  3,
		},
		{
			name:  "testcase with 5 items",
			items: []interface{}{"HH", true, false, 10, 44},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	tests := []struct {
		name  string
		items []interface{}
		want  int
	}{
		{
			name: "empty queue",
			want: 0,
		},
		{
			name:  "5 items queue",
			items: []interface{}{1, 4, 5, "hh", true},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		items []interface{}
		want  bool
	}{
		{
			name: "empty queue",
			want: true,
		},
		{
			name:  "queue with 1 item",
			items: []interface{}{"jj"},
			want:  false,
		},
		{
			name:  "queue with 3 items",
			items: []interface{}{1, 5, 6},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		items   []interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty queue",
			wantErr: true,
		},
		{
			name:  "queue with 1 item",
			items: []interface{}{"Hey"},
			want:  "Hey",
		},
		{
			name:  "queue with 5 items",
			items: []interface{}{1, true, false, "hello", "guy"},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			got, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Multiple_Dequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1).Enqueue(2).Enqueue(3).Enqueue(4).Enqueue(5)
	got, err := q.Dequeue()
	if err != nil {
		t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, false)
		return
	}
	if got != 1 {
		t.Errorf("Queue.Dequeue() = %v, want %v", got, 1)
	}
	got, err = q.Dequeue()
	if err != nil {
		t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, false)
		return
	}
	if got != 2 {
		t.Errorf("Queue.Dequeue() = %v, want %v", got, 2)
	}
	q.Enqueue(30)
	got, err = q.Dequeue()
	if err != nil {
		t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, false)
		return
	}
	if got != 3 {
		t.Errorf("Queue.Dequeue() = %v, want %v", got, 3)
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name    string
		items   []interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty queue",
			wantErr: true,
		},
		{
			name:  "4 items queue",
			items: []interface{}{2, 3, 5, 6},
			want:  2,
		},
		{
			name:  "1 item queue",
			items: []interface{}{4},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			got, err := q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Iterate(t *testing.T) {
	tests := []struct {
		name  string
		items []interface{}
	}{
		{
			name:  "10 items",
			items: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
		{
			name: "empty queue",
		},
		{
			name:  "3 items",
			items: []interface{}{"One", "Two", "Three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			q.Iterate(func(index int, item interface{}) {
				if tt.items[index] != item {
					t.Errorf("index: %d, expected = %v, got = %v", index, tt.items[index], item)
				}
			})
		})
	}
}

func TestQueue_Contains(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name  string
		items []int
		args  args
		want  bool
	}{
		{
			name:  "existing item",
			items: []int{1, 2, 3, 4, 5},
			args: args{
				item: 4,
			},
			want: true,
		},
		{
			name:  "missing item",
			items: []int{1, 2, 3},
			args: args{
				item: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			for _, item := range tt.items {
				q.Enqueue(item)
			}
			if got := q.Contains(tt.args.item); got != tt.want {
				t.Errorf("Stack.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
