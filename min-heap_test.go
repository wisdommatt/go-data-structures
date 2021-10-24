package datastructures

import (
	"reflect"
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	tests := []struct {
		name              string
		items             []float64
		want              []float64
		expectedHashTable map[float64][]int
	}{
		{
			name:  "inserting 3 items",
			items: []float64{3.2, 4, 2},
			want:  []float64{2, 4, 3.2},
			expectedHashTable: map[float64][]int{
				2:   {0},
				3.2: {2},
				4:   {1},
			},
		},
		{
			name:  "inserting 4 items",
			items: []float64{3, 4, 2, 1},
			want:  []float64{1, 2, 3, 4},
			expectedHashTable: map[float64][]int{
				1: {0},
				2: {1},
				3: {2},
				4: {3},
			},
		},
		{
			name:              "inserting no item",
			expectedHashTable: make(map[float64][]int),
		},
		{
			name:  "inserting 1 item",
			items: []float64{6},
			want:  []float64{6},
			expectedHashTable: map[float64][]int{
				6: {0},
			},
		},
		{
			name:  "inserting 10 items",
			items: []float64{9, 4, 6, 2, 6, 3, 7, 8, 3, 10},
			want:  []float64{2, 3, 3, 4, 6, 6, 7, 9, 8, 10},
			expectedHashTable: map[float64][]int{
				2:  {0},
				3:  {2, 1},
				4:  {3},
				6:  {4, 5},
				7:  {6},
				9:  {7},
				8:  {8},
				10: {9},
			},
		},
		{
			name:  "inserting 15 items",
			items: []float64{9, 4, 6, 2, 6, 3, 7, 8, 3, 10, 5, 11, 1, 8, 100},
			want:  []float64{1, 3, 2, 4, 5, 3, 7, 9, 8, 10, 6, 11, 6, 8, 100},
			expectedHashTable: map[float64][]int{
				1:   {0},
				2:   {2},
				3:   {1, 5},
				4:   {3},
				5:   {4},
				6:   {10, 12},
				7:   {6},
				8:   {8, 13},
				9:   {7},
				10:  {9},
				11:  {11},
				100: {14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			if h.Size() != len(tt.want) {
				t.Errorf("Heap.Insert() = %v, want %v", h.Size(), len(tt.want))
			}
			if !reflect.DeepEqual(h.items, tt.want) {
				t.Errorf("Heap.Insert() = %v, want %v", h.items, tt.want)
			}
			if !reflect.DeepEqual(h.hashTable, tt.expectedHashTable) {
				t.Errorf("Heap.Insert() hashTable = %v, \n want %v", h.hashTable, tt.expectedHashTable)
			}
		})
	}
}

func TestMinHeap_Poll(t *testing.T) {
	tests := []struct {
		name              string
		items             []float64
		want              []float64
		pollCount         int
		expectedHashTable map[float64][]int
		wantErr           bool
	}{
		{
			name:      "inserting 4 items - polling 2",
			items:     []float64{3, 5, 2, 6},
			want:      []float64{6},
			pollCount: 3,
			expectedHashTable: map[float64][]int{
				2: {},
				3: {},
				5: {},
				6: {0},
			},
		},
		{
			name:              "inserting 1 item - polling 2",
			items:             []float64{4},
			pollCount:         2,
			wantErr:           true,
			expectedHashTable: map[float64][]int{4: {}},
		},
		{
			name:      "inserting 10 items - polling 9",
			items:     []float64{9, 4, 6, 2, 6, 3, 7, 8, 3, 10},
			want:      []float64{10},
			pollCount: 9,
			expectedHashTable: map[float64][]int{
				2:  {},
				3:  {},
				4:  {},
				6:  {},
				7:  {},
				9:  {},
				8:  {},
				10: {0},
			},
		},
		{
			name:      "inserting 10 items - polling 5",
			items:     []float64{9, 4, 6, 2, 6, 3, 7, 8, 3, 10},
			want:      []float64{6, 8, 7, 10, 9},
			pollCount: 5,
			expectedHashTable: map[float64][]int{
				2:  {},
				3:  {},
				4:  {},
				6:  {0},
				7:  {2},
				9:  {4},
				8:  {1},
				10: {3},
			},
		},
		{
			name:      "inserting 15 items - polling 7",
			items:     []float64{9, 4, 6, 2, 6, 3, 7, 8, 3, 10, 5, 11, 1, 8, 100},
			want:      []float64{6, 8, 7, 9, 10, 8, 11, 100},
			pollCount: 7,
			expectedHashTable: map[float64][]int{
				1:   {},
				2:   {},
				3:   {},
				4:   {},
				5:   {},
				6:   {0},
				7:   {2},
				8:   {5, 1},
				9:   {3},
				10:  {4},
				11:  {6},
				100: {7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			for i := 0; i < tt.pollCount; i++ {
				_, err := h.Poll()
				if err != nil && (err != nil) != tt.wantErr {
					t.Errorf("Heap.Poll() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if h.Size() != len(tt.want) {
				t.Errorf("Heap.Poll() = %v, want %v", h.Size(), len(tt.want))
			}
			if !reflect.DeepEqual(h.items, tt.want) {
				t.Errorf("Heap.Poll() items = %v, \n want %v", h.items, tt.want)
			}
			if !reflect.DeepEqual(h.hashTable, tt.expectedHashTable) {
				t.Errorf("Heap.Poll() hashTable = %v, \n want %v", h.hashTable, tt.expectedHashTable)
			}
		})
	}
}

func TestMinHeap_Remove(t *testing.T) {
	type args struct {
		item float64
	}
	tests := []struct {
		args              args
		name              string
		items             []float64
		want              []float64
		expectedHashTable map[float64][]int
		removedItem       float64
		wantErr           bool
	}{
		{
			name:        "removing first item",
			items:       []float64{20.3, 45.6, 56.2},
			want:        []float64{45.6, 56.2},
			args:        args{item: 20.3},
			removedItem: 20.3,
			expectedHashTable: map[float64][]int{
				20.3: {},
				45.6: {0},
				56.2: {1},
			},
		},
		{
			name:        "removing invalid item",
			items:       []float64{4, 2, 5, 6, 6, 7},
			want:        []float64{2, 4, 5, 6, 6, 7},
			args:        args{item: 3002},
			removedItem: 0,
			expectedHashTable: map[float64][]int{
				2: {0},
				4: {1},
				5: {2},
				6: {3, 4},
				7: {5},
			},
			wantErr: true,
		},
		{
			name:        "removing a middle item",
			items:       []float64{8, 0.33, 0.44, 0.23, 0.12},
			want:        []float64{0.12, 0.23, 0.33, 8},
			args:        args{item: 0.44},
			removedItem: 0.44,
			expectedHashTable: map[float64][]int{
				0.44: {},
				0.12: {0},
				0.23: {1},
				8:    {3},
				0.33: {2},
			},
		},
		{
			name:        "removing an extreme item",
			items:       []float64{1, 5, 9},
			want:        []float64{1, 5},
			args:        args{item: 9},
			removedItem: 9,
			expectedHashTable: map[float64][]int{
				1: {0},
				5: {1},
				9: {},
			},
		},
		{
			name:              "removing from an empty heap",
			args:              args{item: 10},
			wantErr:           true,
			expectedHashTable: map[float64][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			item, err := h.Remove(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Heap.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if item != tt.removedItem {
				t.Errorf("Heap.Remove() item = %v, removedItem %v", item, tt.removedItem)
			}
			if h.Size() != len(tt.want) {
				t.Errorf("Heap.Remove() = %v, want %v", h.Size(), len(tt.want))
			}
			if !reflect.DeepEqual(h.items, tt.want) {
				t.Errorf("Heap.Remove() items = %v, \n want %v", h.items, tt.want)
			}
			if !reflect.DeepEqual(h.hashTable, tt.expectedHashTable) {
				t.Errorf("Heap.Remove() hashTable = %v, \n want %v", h.hashTable, tt.expectedHashTable)
			}
		})
	}
}

func TestMinHeap_Contains(t *testing.T) {
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
			name:  "existing item",
			items: []float64{40, 10, 22, 55},
			args:  args{item: 10},
			want:  true,
		},
		{
			name:  "invalid item",
			items: []float64{30, 49, 233},
			args:  args{item: 443},
		},
		{
			name: "invalid item in empty heap",
			args: args{item: 1344},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			if got := h.Contains(tt.args.item); got != tt.want {
				t.Errorf("Heap.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Size(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  int
	}{
		{
			name: "empty heap",
			want: 0,
		},
		{
			name:  "heap with 6 items",
			items: []float64{2, 4, 5, 6, 2, 5},
			want:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Heap.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Peek(t *testing.T) {
	tests := []struct {
		name    string
		want    float64
		items   []float64
		wantErr bool
	}{
		{
			name:    "peeking from empty heap",
			want:    0,
			wantErr: true,
		},
		{
			name:  "peeking from heap with one item",
			items: []float64{4},
			want:  4,
		},
		{
			name:  "peeking from heep with 5 items",
			items: []float64{8, 4, 5, 3, 6},
			want:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			got, err := h.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("MinHeap.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinHeap.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_GetList(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name: "empty heap",
		},
		{
			name:  "heap with 1 item",
			items: []float64{0},
			want:  []float64{0},
		},
		{
			name:  "heap with 5 items",
			items: []float64{8, 4, 5, 3, 6},
			want:  []float64{3, 4, 5, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeap()
			for _, item := range tt.items {
				h.Insert(item)
			}
			if got := h.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
