package datastructures

import (
	"reflect"
	"testing"
)

func TestAvlTree_Add(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
	}{
		{
			name:  "left-left case",
			items: []float64{5, 4, 3},
		},
		{
			name:  "left-right case",
			items: []float64{5, 3, 4},
		},
		{
			name:  "right-left case",
			items: []float64{3, 5, 4},
		},
		{
			name:  "right-right case",
			items: []float64{3, 4, 5},
		},
		{
			name:  "balanced case",
			items: []float64{4, 5, 3},
		},
		{
			name:  "duplicate balanced case",
			items: []float64{4, 5, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}

			if avl.root.Data != 4 {
				t.Errorf("AvlTree.Add() = %v, want %v", avl.root.Data, 4)
			}
			if avl.root.Right.Data != 5 {
				t.Errorf("AvlTree.Add() = %v, want %v", avl.root.Right.Data, 5)
			}
			if avl.root.Left.Data != 3 {
				t.Errorf("AvlTree.Add() = %v, want %v", avl.root.Left.Data, 3)
			}
		})
	}
}

func TestAvlTree_Search(t *testing.T) {
	type args struct {
		item float64
	}
	tests := []struct {
		name              string
		args              args
		items             []float64
		wantItem          float64
		expectSearchLeft  float64
		expectSearchRight float64
	}{
		{
			name:              "8 items",
			items:             []float64{33, 53, 61, 13, 11, 8, 9, 21},
			args:              args{item: 13},
			wantItem:          13,
			expectSearchLeft:  9,
			expectSearchRight: 53,
		},
		{
			name:              "6 items",
			items:             []float64{6, 9, 10, 3, 2, 5},
			args:              args{item: 3},
			wantItem:          3,
			expectSearchLeft:  2,
			expectSearchRight: 5,
		},
		{
			name:              "11 items",
			items:             []float64{95, 110, 120, 102, 108, 72, 87, 85, 43, 50, 2},
			args:              args{item: 110},
			wantItem:          110,
			expectSearchLeft:  102,
			expectSearchRight: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}
			got := avl.Search(tt.args.item)
			if got != nil && got.Data != tt.wantItem {
				t.Errorf("AvlTree.Search().Data = %v, wantItem %v", got.Data, tt.wantItem)
			}
			if got.Right != nil && got.Right.Data != tt.expectSearchRight {
				t.Errorf("AvlTree.Search().Right = %v, expectSearchRight %v", got.Right.Data, tt.expectSearchRight)
			}
		})
	}

	// testing for invalid node.
	avl := NewAvlTree()
	if got := avl.Search(20); got != nil {
		t.Errorf("AvlTree.Search() want nil got %v", got)
	}
}

// TODO(wisdommatt): add more test cases for larger avl trees to
// test rare cases.
func TestAvlTree_Remove(t *testing.T) {
	avl := NewAvlTree()
	avl.Add(33).Add(53).Add(61).Add(13).Add(11).Add(8).Add(9).Add(21)

	removed := avl.Remove(53)
	if !removed {
		t.Errorf("AvlTree.Remove() want %v, got %v", removed, !removed)
	}
	if avl.GetRoot().Data != 21 {
		t.Errorf("AvlTree.Remove() want %v, got %v", avl.GetRoot().Data, 21)
	}
	if avl.Size() != 7 {
		t.Errorf("AvlTree.Remove() want %v, got %v", avl.Size(), 7)
	}
	if avl.GetRoot().Right.Data != 33 {
		t.Errorf("AvlTree.Remove() want %v, got %v", avl.GetRoot().Right.Data, 33)
	}
	if avl.GetRoot().Left.Data != 9 {
		t.Errorf("AvlTree.Remove() want %v, got %v", avl.GetRoot().Left.Data, 9)
	}
	removed = avl.Remove(9)
	if !removed {
		t.Errorf("AvlTree.Remove() want %v, got %v", true, false)
	}
	if avl.GetRoot().Left.Data != 11 {
		t.Errorf("AvlTree.Remove() want %v, got %v", avl.GetRoot().Left.Data, 11)
	}
	// invalid item.
	removed = avl.Remove(96)
	if removed {
		t.Errorf("AvlTree.Remove() want %v, got %v", false, true)
	}
}

func TestAvlTree_PreOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{3, 1, 0, 2, 5, 4},
		},
		{
			name: "empty tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{5, 3, 1, 0, 4, 9, 7, 6, 8, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}
			var values []float64
			avl.PreOrderTraversal(func(node *AvlNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("AvlTree.PreOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestAvlTree_InOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{0, 1, 2, 3, 4, 5},
		},
		{
			name: "empty tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{0, 1, 3, 4, 5, 6, 7, 8, 9, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}
			var values []float64
			avl.InOrderTraversal(func(node *AvlNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("AvlTree.InOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestAvlTree_PostOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{0, 2, 1, 4, 5, 3},
		},
		{
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{0, 1, 4, 3, 6, 8, 7, 13, 9, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}
			var values []float64
			avl.PostOrderTraversal(func(node *AvlNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("AvlTree.PostOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestAvlTree_LevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{3, 1, 5, 0, 2, 4},
		},
		{
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{5, 3, 9, 1, 4, 7, 13, 0, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avl := NewAvlTree()
			for _, i := range tt.items {
				avl.Add(i)
			}
			var values []float64
			avl.LevelOrderTraversal(func(node *AvlNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("AvlTree.LevelOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}
