package datastructures

import (
	"reflect"
	"testing"
)

func TestHashTable_hash(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "'world' string key",
			args: args{key: "world"},
			want: 6,
		},
		{
			name: "'10' string key",
			args: args{key: "10"},
			want: 8,
		},
		{
			name: "'a' string key",
			args: args{key: "a"},
			want: 0,
		},
		{
			name: "'n' string key",
			args: args{key: "n"},
			want: 3,
		},
		{
			name: "'~~~' string key",
			args: args{key: "~~~"},
			want: 1,
		},
		{
			name: "'~' string key",
			args: args{key: "~"},
			want: 9,
		},
		{
			name: "' ' string key",
			args: args{key: " "},
			want: 5,
		},
		{
			name: "'10' integer key",
			args: args{key: 10},
			want: 4,
		},
		{
			name: "'0' int key",
			args: args{key: 0},
			want: 0,
		},
		{
			name: "'5.2' float32 key",
			args: args{key: float32(5.2)},
			want: 7,
		},
		{
			name: "'5' float32 key",
			args: args{key: float32(5)},
			want: 7,
		},
		{
			name: "'2939948995839849223443349204930940493945943094949493034.4455' float64 key",
			args: args{key: float64(2939948995839849223443349204930940493945943094949493034.4455)},
			want: 0,
		},
		{
			name: "'4448' int16 key",
			args: args{key: int16(4448)},
			want: 4,
		},
		{
			name: "'239993994' int32 key",
			args: args{key: int32(239993994)},
			want: 0,
		},
		{
			name: "'399948959896899385993899' int64 key",
			args: args{key: int64(3999489598968993859)},
			want: 7,
		},
		{
			name:    "invalid key",
			args:    args{key: struct{}{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(10)
			got, err := h.hash(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashTable.hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashTable.hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Set(t *testing.T) {
	type args struct {
		key   interface{}
		value interface{}
	}
	testCase := struct {
		name             string
		items            []args
		expectedElements int
	}{
		name: "key with 0 hash",
		items: []args{
			{key: "world", value: "the world"},
			{key: "world", value: "updated world"},
			{key: "fish", value: "fish world"},
			{key: "make", value: "make world"},
		},
		expectedElements: 3,
	}
	t.Run(testCase.name, func(t *testing.T) {
		h := NewHashTable(2)
		for _, item := range testCase.items {
			h.Set(item.key, item.value)
		}
		if h.Elements() != testCase.expectedElements {
			t.Errorf("HashTable.Set() = %v, want %v", h.Elements(), testCase.expectedElements)
		}
		headValue := h.table[0].GetHead().Data.(HashTableEntry).Value
		if headValue != "updated world" {
			t.Errorf("HashTable.Set() = %v, want %v", headValue, "updated world")
		}
		secondValue := h.table[1].GetHead().Data.(HashTableEntry).Value
		if secondValue != "make world" {
			t.Errorf("HashTable.Set() = %v, want %v", secondValue, "make world")
		}
		headNextValue := h.table[0].GetHead().Next.Data.(HashTableEntry).Value
		if headNextValue != "fish world" {
			t.Errorf("HashTable.Set() = %v, want %v", headNextValue, "fish world")
		}
	})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "invalid key",
			args:    args{key: []int{23}, value: "value"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(20)
			if err := h.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("HashTable.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type args struct {
		key interface{}
	}
	type item struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		items   []item
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "existing key",
			items: []item{
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
			args: args{key: 4},
			want: "4th",
		},
		{
			name: "existing string key",
			items: []item{
				{key: "1 key", value: "1st"},
				{key: "2 key", value: "2nd"},
				{key: "3 key", value: "3rd"},
				{key: "4 key", value: "4th"},
				{key: "5 key", value: "5th"},
			},
			args: args{key: "3 key"},
			want: "3rd",
		},
		{
			name: "non-existing key",
			items: []item{
				{key: "1 key", value: "1st"},
				{key: "2 key", value: "2nd"},
				{key: "3 key", value: "3rd"},
				{key: "4 key", value: "4th"},
				{key: "5 key", value: "5th"},
			},
			args:    args{key: "invalid key"},
			wantErr: true,
		},
		{
			name: "non-existing key with existing key-hash",
			items: []item{
				{key: 0, value: "0th"},
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
			args:    args{key: 6},
			wantErr: true,
		},
		{
			name:    "empty hash-table",
			args:    args{key: "kk"},
			wantErr: true,
		},
		{
			name:    "invalid key",
			args:    args{key: nil},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(2)
			for _, item := range tt.items {
				h.Set(item.key, item.value)
			}
			got, err := h.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashTable.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashTable.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	type args struct {
		key interface{}
	}
	type item struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name             string
		args             args
		items            []item
		expectedElements int
	}{
		{
			name: "existing key",
			items: []item{
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
			args:             args{key: 4},
			expectedElements: 4,
		},
		{
			name: "existing string key",
			items: []item{
				{key: "1 key", value: "1st"},
				{key: "2 key", value: "2nd"},
				{key: "3 key", value: "3rd"},
				{key: "4 key", value: "4th"},
				{key: "5 key", value: "5th"},
			},
			args:             args{key: "3 key"},
			expectedElements: 4,
		},
		{
			name: "non-existing key",
			items: []item{
				{key: "1 key", value: "1st"},
				{key: "2 key", value: "2nd"},
				{key: "3 key", value: "3rd"},
				{key: "4 key", value: "4th"},
				{key: "5 key", value: "5th"},
			},
			args:             args{key: "invalid key"},
			expectedElements: 5,
		},
		{
			name: "non-existing key with existing key-hash",
			items: []item{
				{key: 0, value: "0th"},
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
			args:             args{key: 6},
			expectedElements: 6,
		},
		{
			name:             "empty hash-table",
			args:             args{key: "kk"},
			expectedElements: 0,
		},
		{
			name: "invalid key",
			items: []item{
				{key: 0, value: "0th"},
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
			args:             args{key: nil},
			expectedElements: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(2)
			for _, item := range tt.items {
				h.Set(item.key, item.value)
			}
			h.Delete(tt.args.key)
			if h.Elements() != tt.expectedElements {
				t.Errorf("HashTable.Set() = %v, want %v", h.Elements(), tt.expectedElements)
			}
		})
	}
}

func TestHashTable_Size(t *testing.T) {
	type fields struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "size 14",
			fields: fields{size: 14},
			want:   14,
		},
		{
			name:   "size 0",
			fields: fields{size: 0},
			want:   0,
		},
		{
			name:   "size 3",
			fields: fields{size: 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(tt.fields.size)
			if got := h.Size(); got != tt.want {
				t.Errorf("HashTable.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Iterate(t *testing.T) {
	type item struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name  string
		items []item
	}{
		{
			name: "6 items",
			items: []item{
				{key: 0, value: "0th"},
				{key: 1, value: "1st"},
				{key: 2, value: "2nd"},
				{key: 3, value: "3rd"},
				{key: 4, value: "4th"},
				{key: 5, value: "5th"},
			},
		},
		{
			name:  "0 items",
			items: []item{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(20)
			for _, item := range tt.items {
				h.Set(item.key, item.value)
			}
			result := []item{}
			h.Iterate(func(key, value interface{}) {
				result = append(result, item{key, value})
			})
			if !reflect.DeepEqual(tt.items, result) {
				t.Errorf("HashTable.Iterate() = %v, want %v", tt.items, result)
			}
		})
	}
}
