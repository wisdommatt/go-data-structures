package datastructures

// BstNode is the node used in the binary search tree data
// structure.
type BstNode struct {
	Left  *BstNode
	Right *BstNode
	Data  float64
}

// BinarySearchTree represents the binary search tree data
// structure.
//
// this binary tree search does not support duplicate items.
type BinarySearchTree struct {
	root *BstNode
	size int
}

// NewBinarySearchTree returns a new binary search tree data
// structure.
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Add adds an item to the binary search tree.
func (b *BinarySearchTree) Add(elem float64) *BinarySearchTree {
	if b.Size() == 0 {
		b.root = &BstNode{Data: elem}
		b.size++
		return b
	}
	if b.find(b.root, elem) != nil {
		return b
	}
	b.root = b.insert(b.root, elem)
	b.size++
	return b
}

func (b *BinarySearchTree) insert(node *BstNode, elem float64) *BstNode {
	if node == nil {
		return &BstNode{Data: elem}
	}
	if elem > node.Data {
		node.Right = b.insert(node.Right, elem)
	} else {
		node.Left = b.insert(node.Left, elem)
	}
	return node
}

// Search walks through the binary search tree to look for
// the specified item.
func (b *BinarySearchTree) Search(elem float64) *BstNode {
	return b.find(b.root, elem)
}

func (b *BinarySearchTree) find(node *BstNode, elem float64) *BstNode {
	var item *BstNode
	if node == nil {
		return item
	}
	if elem == node.Data {
		item = node
	}
	if elem > node.Data {
		item = b.find(node.Right, elem)
	}
	if elem < node.Data {
		item = b.find(node.Left, elem)
	}
	return item
}

// Size returns the size of the binary search tree.
func (b *BinarySearchTree) Size() int {
	return b.size
}

// Remove removes an item from the binary search tree.
func (b *BinarySearchTree) Remove(elem float64) bool {
	if b.find(b.root, elem) == nil {
		return false
	}
	b.root = b.removeItem(b.GetRoot(), elem)
	b.size--
	return true
}

func (b *BinarySearchTree) removeItem(node *BstNode, elem float64) *BstNode {
	if elem > node.Data {
		node.Right = b.removeItem(node.Right, elem)
		return node
	}
	if elem < node.Data {
		node.Left = b.removeItem(node.Left, elem)
		return node
	}
	// replacing/swaping the deleted node with the left child if
	// right child is nil.
	if node.Left != nil && node.Right == nil {
		return node.Left
	}
	// replacing/swaping the deleted node with right child if
	// left child is nil.
	if node.Right != nil && node.Left == nil {
		return node.Right
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	// replacing the deleted node with smallest node in the right
	// subtree if both left and right subtrees exist.
	//
	// this is done to balance the binary search tree invariant.
	replacement := b.findMinNode(node.Right)
	node.Data = replacement.Data
	node.Right = b.removeItem(node.Right, replacement.Data)
	return node
}

// findMinNode is a helper function for retrieving the smallest node
// in a subtree.
func (b *BinarySearchTree) findMinNode(node *BstNode) *BstNode {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// GetRoot returns the root node of the binary search tree.
func (b *BinarySearchTree) GetRoot() *BstNode {
	return b.root
}

// PreOrderTraversal runs a pre-order traversal on the binary search
// tree and execute the callback function f for each iteration.
//
// pre-order traversal follows the order: <root>-<left>-<right>
func (b *BinarySearchTree) PreOrderTraversal(f func(node *BstNode)) {
	b.runPreOrderTraversal(b.root, f)
}

func (b *BinarySearchTree) runPreOrderTraversal(node *BstNode, f func(node *BstNode)) {
	if node == nil {
		return
	}
	f(node)
	b.runPreOrderTraversal(node.Left, f)
	b.runPreOrderTraversal(node.Right, f)
}

// InOrderTraversal runs an in-order traversal on the binary search
// tree and execute the callback function f for each iteration.
//
// in-order traversal follows the order: <left>-<root>-<right>
func (b *BinarySearchTree) InOrderTraversal(f func(node *BstNode)) {
	b.runInOrderTraversal(b.root, f)
}

func (b *BinarySearchTree) runInOrderTraversal(node *BstNode, f func(node *BstNode)) {
	if node == nil {
		return
	}
	b.runInOrderTraversal(node.Left, f)
	f(node)
	b.runInOrderTraversal(node.Right, f)
}

// PostOrderTraversal runs a post-order traversal on the binary search
// tree and execute the callback function f for each iteration.
//
// post-order traversal follows the order: <left>-<right>-<root>
func (b *BinarySearchTree) PostOrderTraversal(f func(node *BstNode)) {
	b.runPostOrderTraversal(b.root, f)
}

func (b *BinarySearchTree) runPostOrderTraversal(node *BstNode, f func(node *BstNode)) {
	if node == nil {
		return
	}
	b.runPostOrderTraversal(node.Left, f)
	b.runPostOrderTraversal(node.Right, f)
	f(node)
}

// LevelOrderTraversal runs a level-order traversal on the binary search
// tree and execute the callback function f for each iteration.
func (b *BinarySearchTree) LevelOrderTraversal(f func(node *BstNode)) {
	b.runLevelOrderTraversal(b.root, f)
}

func (b *BinarySearchTree) runLevelOrderTraversal(node *BstNode, f func(node *BstNode)) {
	if node == nil {
		return
	}
	queue := NewQueue()
	queue.Enqueue(node)
	for !queue.IsEmpty() {
		qNodeI, _ := queue.Dequeue()
		qNode := qNodeI.(*BstNode)
		f(qNode)
		if qNode.Left != nil {
			queue.Enqueue(qNode.Left)
		}
		if qNode.Right != nil {
			queue.Enqueue(qNode.Right)
		}
	}
}
