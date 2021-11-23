package datastructures

import (
	"reflect"
	"testing"
)

func TestNewFenwickTree(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *FenwickTree
	}{
		{
			name: "12 items",
			args: args{data: []int{3, 4, -2, 7, 3, 11, 5, -8, -9, 2, 4, -8}},
			want: &FenwickTree{
				tree: []int{0, 3, 7, -2, 12, 3, 14, 5, 23, -9, -7, 4, -11},
			},
		},
		{
			name: "random list 1",
			args: args{data: []int{9, 91, 10, 92, 19, 50, 95, 31, 90, 23, 98}},
			want: &FenwickTree{
				tree: []int{0, 9, 100, 10, 202, 19, 69, 95, 397, 90, 113, 98},
			},
		},
		{
			name: "random list 2",
			args: args{data: []int{54, 45, 58, 4, 65}},
			want: &FenwickTree{
				tree: []int{0, 54, 99, 58, 161, 65},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFenwickTree(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFenwickTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFenwickTree_PrefixSum(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "randomg data 1",
			args:   args{i: 3},
			fields: fields{data: []int{3, 4, -2, 7, 3, 11, 5, -8, -9, 2, 4, -8}},
			want:   5,
		},
		{
			name:   "randomg data 2",
			args:   args{i: 5},
			fields: fields{data: []int{9, 91, 10, 92, 19, 50, 95, 31, 90, 23, 98}},
			want:   221,
		},
		{
			name:   "last index",
			args:   args{i: 5},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   226,
		},
		{
			name:   "first index",
			args:   args{i: 1},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   54,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := NewFenwickTree(tt.fields.data)
			if got := ft.PrefixSum(tt.args.i); got != tt.want {
				t.Errorf("FenwickTree.PrefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFenwickTree_RangeQuery(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "randomg data 1",
			args:   args{j: 3, i: 2},
			fields: fields{data: []int{3, 4, -2, 7, 3, 11, 5, -8, -9, 2, 4, -8}},
			want:   2,
		},
		{
			name:   "randomg data 2",
			args:   args{j: 5, i: 3},
			fields: fields{data: []int{9, 91, 10, 92, 19, 50, 95, 31, 90, 23, 98}},
			want:   121,
		},
		{
			name:   "equal i and j",
			args:   args{j: 5, i: 5},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   65,
		},
		{
			name:   "i greater than j",
			args:   args{i: 5, j: 3},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := NewFenwickTree(tt.fields.data)
			if got := ft.RangeQuery(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("FenwickTree.RangeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFenwickTree_PointAdd(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		p   int
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "randomg data 1",
			args:   args{p: 3, val: 12},
			fields: fields{data: []int{3, 4, -2, 7, 3, 11, 5, -8, -9, 2, 4, -8}},
			want:   []int{0, 3, 7, 10, 24, 3, 14, 5, 35, -9, -7, 4, -11},
		},
		{
			name:   "randomg data 2",
			args:   args{p: 5, val: 8},
			fields: fields{data: []int{9, 91, 10, 92, 19, 50, 95, 31, 90, 23, 98}},
			want:   []int{0, 9, 100, 10, 202, 27, 77, 95, 405, 90, 113, 98},
		},
		{
			name:   "last index",
			args:   args{p: 5, val: 300},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   []int{0, 54, 99, 58, 161, 365},
		},
		{
			name:   "first index",
			args:   args{p: 1, val: 222},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   []int{0, 276, 321, 58, 383, 65},
		},
		{
			name:   "non-existing index",
			args:   args{p: 10, val: 111},
			fields: fields{data: []int{54, 45, 58, 4, 65}},
			want:   []int{0, 54, 99, 58, 161, 65},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := NewFenwickTree(tt.fields.data)
			ft.PointAdd(tt.args.p, tt.args.val)
			if !reflect.DeepEqual(ft.tree, tt.want) {
				t.Errorf("NewFenwickTree() = %v, want %v", ft.tree, tt.want)
			}
		})
	}
}
