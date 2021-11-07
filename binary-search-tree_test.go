package datastructures

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_Add(t *testing.T) {
	tests := []struct {
		name            string
		items           []float64
		expectRootLeft  *BstNode
		expectRootRight *BstNode
	}{
		{
			name:            "6 items",
			items:           []float64{2, 1, 0, 3, 5, 4},
			expectRootRight: &BstNode{Data: 3},
			expectRootLeft:  &BstNode{Data: 1},
		},
		{
			name:            "1 item",
			items:           []float64{4},
			expectRootLeft:  nil,
			expectRootRight: nil,
		},
		{
			name:            "duplicate items",
			items:           []float64{2, 1, 0, 3, 5, 4, 1},
			expectRootRight: &BstNode{Data: 3},
			expectRootLeft:  &BstNode{Data: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			if got := b.root.Right; got != nil && got.Data != tt.expectRootRight.Data {
				t.Errorf("BinarySearchTree.root.Right = %v, expectRootRight %v", got.Data, tt.expectRootRight.Data)
			}
			if got := b.root.Left; got != nil && got.Data != tt.expectRootLeft.Data {
				t.Errorf("BinarySearchTree.root.Left = %v, expectRootLeft %v", got.Data, tt.expectRootLeft.Data)
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	type args struct {
		elem float64
	}
	tests := []struct {
		name              string
		args              args
		want              *BstNode
		items             []float64
		wantItem          float64
		expectSearchLeft  float64
		expectSearchRight float64
	}{
		{
			name:              "7 items",
			items:             []float64{2, 1, 0, 3, 2.5, 5, 4},
			wantItem:          3,
			args:              args{elem: 3},
			expectSearchLeft:  2.5,
			expectSearchRight: 5,
		},
		{
			name:     "1 item",
			items:    []float64{5},
			wantItem: 5,
			args:     args{5},
		},
		{
			name:     "8 items",
			items:    []float64{2, 1, 0, 3, 2.5, 5, 4, 9},
			wantItem: 0,
			args:     args{elem: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			got := b.Search(tt.args.elem)
			if got != nil && got.Data != tt.wantItem {
				t.Errorf("BinarySearchTree.Search().Data = %v, wantItem %v", got.Data, tt.wantItem)
			}
			if got.Right != nil && got.Right.Data != tt.expectSearchRight {
				t.Errorf("BinarySearchTree.Search().Right = %v, expectSearchRight %v", got.Right.Data, tt.expectSearchRight)
			}
			if got.Left != nil && got.Left.Data != tt.expectSearchLeft {
				t.Errorf("BinarySearchTree.Search().Left = %v, expectSearchLeft %v", got.Left.Data, tt.expectSearchLeft)
			}
		})
	}
}

func TestBinarySearchTree_Remove(t *testing.T) {
	type args struct {
		elem float64
	}
	tests := []struct {
		name            string
		args            args
		want            bool
		items           []float64
		expectRootLeft  *BstNode
		expectRootRight *BstNode
		expectSize      int
	}{
		{
			name: "non-existing item",
			args: args{elem: 12},
			want: false,
		},
		{
			name:            "removing the root node",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 4},
			args:            args{2},
			want:            true,
			expectRootLeft:  &BstNode{Data: 1},
			expectRootRight: &BstNode{Data: 3},
			expectSize:      6,
		},
		{
			name:            "removing a node greater than the root node",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 4},
			args:            args{3},
			want:            true,
			expectRootLeft:  &BstNode{Data: 1},
			expectRootRight: &BstNode{Data: 4},
			expectSize:      6,
		},
		{
			name:            "removing a node less than the root node",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 4},
			args:            args{1},
			want:            true,
			expectRootLeft:  &BstNode{Data: 0},
			expectRootRight: &BstNode{Data: 3},
			expectSize:      6,
		},
		{
			name:            "removing a left leaf node",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 4},
			args:            args{0},
			want:            true,
			expectRootLeft:  &BstNode{Data: 1},
			expectRootRight: &BstNode{Data: 3},
			expectSize:      6,
		},
		{
			name:            "removing a right leaf node",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 4},
			args:            args{4},
			want:            true,
			expectRootLeft:  &BstNode{Data: 1},
			expectRootRight: &BstNode{Data: 3},
			expectSize:      6,
		},
		{
			name:            "removing a node with null left subtree",
			items:           []float64{2, 1, 0, 3, 2.4, 5, 7},
			args:            args{5},
			want:            true,
			expectRootLeft:  &BstNode{Data: 1},
			expectRootRight: &BstNode{Data: 3},
			expectSize:      6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			if got := b.Remove(tt.args.elem); got != tt.want {
				t.Errorf("BinarySearchTree.Remove() = %v, want %v", got, tt.want)
			}
			if b.root != nil {
				if got := b.root.Right; got != nil && got.Data != tt.expectRootRight.Data {
					t.Errorf("BinarySearchTree.root.Right = %v, expectRootRight %v", got.Data, tt.expectRootRight.Data)
				}
				if got := b.root.Left; got != nil && got.Data != tt.expectRootLeft.Data {
					t.Errorf("BinarySearchTree.root.Left = %v, expectRootLeft %v", got.Data, tt.expectRootLeft.Data)
				}
			}
			if b.Size() != tt.expectSize {
				t.Errorf("BinarySearchTree.Size() = %v, expectSize %v", b.Size(), tt.expectSize)
			}
		})
	}
}

func TestBinarySearchTree_PreOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{2, 1, 0, 3, 5, 4},
		},
		{
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{9, 3, 1, 0, 5, 4, 7, 6, 8, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			var values []float64
			b.PreOrderTraversal(func(node *BstNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("BinarySearchTree.PreOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_InOrderTraversal(t *testing.T) {
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
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{0, 1, 3, 4, 5, 6, 7, 8, 9, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			var values []float64
			b.InOrderTraversal(func(node *BstNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("BinarySearchTree.InOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_PostOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{0, 1, 4, 5, 3, 2},
		},
		{
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{0, 1, 4, 6, 8, 7, 5, 3, 13, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			var values []float64
			b.PostOrderTraversal(func(node *BstNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("BinarySearchTree.PostOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_LevelOrderTraversal(t *testing.T) {
	tests := []struct {
		name  string
		items []float64
		want  []float64
	}{
		{
			name:  "6 items",
			items: []float64{2, 1, 0, 3, 5, 4},
			want:  []float64{2, 1, 3, 0, 5, 4},
		},
		{
			name: "empty binary search tree",
		},
		{
			name:  "10 items",
			items: []float64{9, 3, 5, 1, 4, 7, 13, 0, 6, 8},
			want:  []float64{9, 3, 13, 1, 5, 0, 4, 7, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBinarySearchTree()
			for _, i := range tt.items {
				b.Add(i)
			}
			var values []float64
			b.LevelOrderTraversal(func(node *BstNode) {
				values = append(values, node.Data)
			})
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("BinarySearchTree.LevelOrderTraversal() = %v, want %v", values, tt.want)
			}
		})
	}
}
