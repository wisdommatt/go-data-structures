package datastructures

import (
	"math"
	"math/bits"
)

// FenwickTree represents a fenwick tree data structure.
type FenwickTree struct {
	tree []int
}

// NewFenwickTree returns a new fenwick tree data structure.
func NewFenwickTree(data []int) *FenwickTree {
	length := len(data) + 1 // index starts from 1
	ft := &FenwickTree{
		tree: make([]int, length),
	}
	for k, v := range data {
		ft.tree[k+1] = v
	}
	for k, v := range ft.tree {
		if k == 0 { // excluding index 0
			continue
		}
		j := k + ft.lsb(k)
		if j < length {
			ft.tree[j] = ft.tree[j] + v
		}
	}
	return ft
}

// PrefixSum returns the prefix sum for i.
func (ft *FenwickTree) PrefixSum(index int) int {
	sum := 0
	for index != 0 {
		sum += ft.tree[index]
		index = index - ft.lsb(index)
	}
	return sum
}

// RangeQuery performs a range query on the fenwick tree and
// return the value.
func (ft *FenwickTree) RangeQuery(i int, j int) int {
	return ft.PrefixSum(j) - ft.PrefixSum(i-1)
}

// PointAdd updates items in the fenwick tree starting from the
// point p.
func (ft *FenwickTree) PointAdd(p int, val int) {
	for p < ft.Size() {
		ft.tree[p] = ft.tree[p] + val
		p = p + ft.lsb(p)
	}
}

// Size returns the size of the fenwick tree.
//
// the fenwick tree is one index larger than the array
// it was created from because indexing in the fenwick
// tree starts from 1 not 0.
func (ft *FenwickTree) Size() int {
	return len(ft.tree)
}

// lsb is a helper function for retrieving the least significant
// bit of an integer.
func (ft *FenwickTree) lsb(number int) int {
	trailingZeros := bits.TrailingZeros(uint(number))
	if trailingZeros < 1 {
		return 1
	}
	return int(math.Pow(2, float64(trailingZeros)))
}
