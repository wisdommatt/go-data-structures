package datastructures

import (
	"reflect"
	"testing"
)

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "empty test case",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			if got := s.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "empty test case",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
		{
			name: "sample test case",
			args: args{
				item: "Hello World",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			s.Push(tt.args.item)
			if got := s.Size(); got != 1 {
				t.Errorf("Stack.Size() = %v, want %v", s.Size(), 1)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		items   []interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:  "3 items",
			items: []interface{}{"Hello", "World", "Guys"},
			want:  "Guys",
		},
		{
			name:    "empty stack",
			wantErr: true,
		},
		{
			name:  "1 item stack",
			items: []interface{}{"Hello"},
			want:  "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, item := range tt.items {
				s.Push(item)
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Multiple_Pop(t *testing.T) {
	s := NewStack()
	s.Push("One").Push("Two").Push("Three").Push("Four").Push("Five")
	poped, err := s.Pop()
	if err != nil {
		t.Errorf("Stack.Pop() error = %v, wantErr %v", err, false)
		return
	}
	if poped != "Five" {
		t.Errorf("Stack.Pop() = %v, want %v", poped, "Five")
	}
	poped, err = s.Pop()
	if err != nil {
		t.Errorf("Stack.Pop() error = %v, wantErr %v", err, false)
		return
	}
	if poped != "Four" {
		t.Errorf("Stack.Pop() = %v, want %v", poped, "Four")
	}
	poped, err = s.Pop()
	if err != nil {
		t.Errorf("Stack.Pop() error = %v, wantErr %v", err, false)
		return
	}
	if poped != "Three" {
		t.Errorf("Stack.Pop() = %v, want %v", poped, "Three")
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name    string
		items   []string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "empty stack",
			wantErr: true,
		},
		{
			name:  "5 items stack",
			items: []string{"One", "Two", "Three", "Four", "Five"},
			want:  "Five",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			for _, item := range tt.items {
				s.Push(item)
			}
			got, err := s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Iterate(t *testing.T) {
	tests := []struct {
		name  string
		items []string
	}{
		{
			name:  "10 items",
			items: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		},
		{
			name: "empty stack",
		},
		{
			name:  "3 items",
			items: []string{"One", "Two", "Three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack()
			s.Iterate(func(index int, item interface{}) {
				if tt.items[index] != item {
					t.Errorf("index: %d, expected = %v, got = %v", index, tt.items[index], item)
				}
			})
		})
	}
}

func TestStack_Contains(t *testing.T) {
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
			s := NewStack()
			for _, item := range tt.items {
				s.Push(item)
			}
			if got := s.Contains(tt.args.item); got != tt.want {
				t.Errorf("Stack.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
