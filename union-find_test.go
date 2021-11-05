package datastructures

import (
	"testing"
)

func TestUnionFind_Find(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name               string
		args               args
		elements           int
		want               int
		union              [][]int
		expectedComponents int
	}{
		{
			name: "15 items",
			args: args{i: 6},
			union: [][]int{
				{0, 3}, {0, 10}, {11, 12}, {6, 8}, {11, 0}, {12, 8}, {10, 3},
			},
			elements:           15,
			want:               3,
			expectedComponents: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUnionFind(tt.elements)
			for _, unions := range tt.union {
				u.Unify(unions[0], unions[1])
			}
			if got := u.Find(tt.args.i); got != tt.want {
				t.Errorf("UnionFind.Find() = %v, want %v", got, tt.want)
			}
			if got := u.Components(); got != tt.expectedComponents {
				t.Errorf("UnionFind.Components() = %v, expectedComponents %v", got, tt.expectedComponents)
			}
		})
	}
}

func TestUnionFind_Size(t *testing.T) {
	tests := []struct {
		name     string
		elements int
		want     int
	}{
		{
			name:     "100 items",
			elements: 100,
			want:     100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUnionFind(tt.elements)
			if got := u.Size(); got != tt.want {
				t.Errorf("UnionFind.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
