package datastructures

import (
	"reflect"
	"testing"
)

func TestNewSuffixArray(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want *SuffixArray
	}{
		{
			name: "hello",
			args: args{text: "hello"},
			want: &SuffixArray{
				text:                 "hello",
				size:                 5,
				suffixArr:            []int{1, 0, 2, 3, 4},
				hasComputedSuffixArr: true,
				hasComputedLCPArr:    true,
				lcpArr:               []int{0, 0, 0, 1, 0},
			},
		},
		{
			name: "mississippi",
			args: args{text: "mississippi"},
			want: &SuffixArray{
				text:                 "mississippi",
				size:                 11,
				suffixArr:            []int{10, 7, 4, 1, 0, 9, 8, 6, 3, 5, 2},
				hasComputedSuffixArr: true,
				hasComputedLCPArr:    true,
				lcpArr:               []int{0, 1, 1, 4, 0, 0, 1, 0, 2, 1, 3},
			},
		},
		{
			name: "Wolloomooloo",
			args: args{text: "Wolloomooloo"},
			want: &SuffixArray{
				text:                 "Wolloomooloo",
				size:                 12,
				suffixArr:            []int{0, 2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4},
				hasComputedSuffixArr: true,
				hasComputedLCPArr:    true,
				lcpArr:               []int{0, 0, 1, 3, 0, 0, 1, 2, 1, 1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSuffixArray(tt.args.text)
			got.buildSuffixArray()
			got.buildLCPArray()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSuffixArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuffixArray_GetSuffixArray(t *testing.T) {
	tests := []struct {
		name string
		sf   *SuffixArray
		want []int
	}{
		{
			name: "computed suffix array",
			sf: &SuffixArray{
				text:                 "Wolloomooloo",
				size:                 12,
				suffixArr:            []int{0, 2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4},
				hasComputedSuffixArr: true,
			},
			want: []int{0, 2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4},
		},
		{
			name: "not-computed suffix array",
			sf: &SuffixArray{
				text: "Wolloomooloo",
			},
			want: []int{0, 2, 9, 3, 6, 11, 1, 8, 5, 10, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sf.GetSuffixArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuffixArray.GetSuffixArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuffixArray_UniqueSubstrings(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "AZAZA text",
			fields: fields{text: "AZAZA"},
			want:   9,
		},
		{
			name:   "hello text",
			fields: fields{text: "hello"},
			want:   14,
		},
		{
			name:   "'wow guys' text",
			fields: fields{text: "wow guys"},
			want:   35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := NewSuffixArray(tt.fields.text)
			if got := sa.UniqueSubstrings(); got != tt.want {
				t.Errorf("SuffixArray.UniqueSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuffixArray_GetLCPArray(t *testing.T) {
	tests := []struct {
		name string
		sf   *SuffixArray
		want []int
	}{
		{
			name: "computed LCP array",
			sf: &SuffixArray{
				text:              "Wolloomooloo",
				size:              12,
				lcpArr:            []int{0, 0, 1, 3, 0, 0, 1, 2, 1, 1, 2, 2},
				hasComputedLCPArr: true,
			},
			want: []int{0, 0, 1, 3, 0, 0, 1, 2, 1, 1, 2, 2},
		},
		{
			name: "not-computed LCP array",
			sf: &SuffixArray{
				text: "Wolloomooloo",
				size: 12,
			},
			want: []int{0, 0, 1, 3, 0, 0, 1, 2, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sf.GetLCPArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuffixArray.GetLCPArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
