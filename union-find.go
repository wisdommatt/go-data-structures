package datastructures

// UnionFind represents the union find data structure.
type UnionFind struct {
	sizes           []int
	ids             []int
	size            int
	componentsCount int
}

// NewUnionFind returns a new union find data structure object.
func NewUnionFind(size int) *UnionFind {
	unionFind := &UnionFind{
		size:            size,
		componentsCount: size,
	}
	for i := 0; i < size; i++ {
		unionFind.ids = append(unionFind.ids, i)
		unionFind.sizes = append(unionFind.sizes, 1)
	}
	return unionFind
}

// Find returns the root/component representative of an item.
func (u *UnionFind) Find(i int) int {
	root := i
	for root != u.ids[root] {
		root = u.ids[root]
	}

	// compressing the component path to reduce the time
	// required to retrieve the root of any of the components.
	for root != u.ids[i] {
		next := u.ids[i]
		u.ids[i] = root
		i = next
	}
	return root
}

// Unify connects two items together.
func (u *UnionFind) Unify(i1 int, i2 int) {
	if u.IsConnected(i1, i2) {
		return
	}
	rootI1 := u.Find(i1)
	rootI2 := u.Find(i2)
	if u.sizes[rootI2] < u.sizes[rootI1] {
		u.ids[rootI2] = rootI1
		u.sizes[rootI1] += u.sizes[rootI2]
	} else {
		u.ids[rootI1] = rootI2
		u.sizes[rootI2] += u.sizes[rootI1]
	}
	u.componentsCount--
}

// IsConnected returns true if i1 and i2 are connected.
func (u *UnionFind) IsConnected(i1 int, i2 int) bool {
	rootI1 := u.Find(i1)
	rootI2 := u.Find(i2)
	return rootI1 == rootI2
}

// Size returns the number of items in the data structure.
func (u *UnionFind) Size() int {
	return u.size
}

// Components returns the number of components in the union
// find data structure.
func (u *UnionFind) Components() int {
	return u.componentsCount
}
