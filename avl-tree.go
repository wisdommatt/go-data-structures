package datastructures

import (
	"math"
)

// AvlTree represents an avl tree data structure.
type AvlTree struct {
	root *AvlNode
	size int
}

// AvlNode is the node used in the avl tree data structure.
type AvlNode struct {
	bf     int
	height int
	Left   *AvlNode
	Right  *AvlNode
	Data   float64
}

// NewAvlTree returns a new avl tree data structure.
func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

// Add adds a new node to the avl tree.
func (avl *AvlTree) Add(elem float64) *AvlTree {
	avl.root = avl.insert(avl.root, elem)
	avl.size++
	return avl
}

// insert is a helper method for inserting new nodes in the avl
// tree.
func (avl *AvlTree) insert(node *AvlNode, elem float64) *AvlNode {
	if node == nil {
		return &AvlNode{Data: elem}
	}
	if elem == node.Data {
		return node // do nothing
	}
	if elem > node.Data {
		node.Right = avl.insert(node.Right, elem)
	}
	if elem < node.Data {
		node.Left = avl.insert(node.Left, elem)
	}
	avl.update(node)
	return avl.balance(node)
}

// update is a helper method to update the height and balance factor
// of a node.
func (avl *AvlTree) update(node *AvlNode) {
	leftHeight := -1
	rightHeight := -1
	if node.Left != nil {
		leftHeight = node.Left.height
	}
	if node.Right != nil {
		rightHeight = node.Right.height
	}
	node.bf = rightHeight - leftHeight
	node.height = 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

// balance is the helper method to balance the avl tree if the balance
// factor not in {-1, 0, +1}.
func (avl *AvlTree) balance(node *AvlNode) *AvlNode {
	// left heavy
	if node.bf == -2 {
		if node.Left.bf <= 0 {
			return avl.leftLeftCaseRotation(node)
		}
		return avl.leftRightCaseRotation(node)
	}

	// right heavy
	if node.bf == 2 {
		if node.Right.bf >= 0 {
			return avl.rightRightCaseRotation(node)
		}
		return avl.rightLeftCaseRotation(node)
	}
	return node
}

// leftLeftCaseRotation is a helper method to handle left-left
// case rotation.
func (avl *AvlTree) leftLeftCaseRotation(node *AvlNode) *AvlNode {
	return avl.rotateRight(node)
}

// leftRightCaseRotation is a helper method to handle left-right
// case rotation.
func (avl *AvlTree) leftRightCaseRotation(node *AvlNode) *AvlNode {
	node.Left = avl.rotateLeft(node.Left)
	return avl.leftLeftCaseRotation(node)
}

// rightLeftCaseRotation is a helper method to handle right-left
// case rotation.
func (avl *AvlTree) rightLeftCaseRotation(node *AvlNode) *AvlNode {
	node.Right = avl.rotateRight(node.Right)
	return avl.rightRightCaseRotation(node)
}

// rightRightCaseRotation is a helper method to handle  right-right
// case rotation.
func (avl *AvlTree) rightRightCaseRotation(node *AvlNode) *AvlNode {
	return avl.rotateLeft(node)
}

// rotateRight is a helper method to do a right rotation on
// a node.
func (avl *AvlTree) rotateRight(node *AvlNode) *AvlNode {
	leftNode := node.Left
	node.Left = leftNode.Right
	leftNode.Right = node
	// updating the height and balance factor for the rotated
	// nodes.
	avl.update(node)
	avl.update(leftNode)
	return leftNode
}

// rotateLeft is a helper method to do a left rotation on
// a node.
func (avl *AvlTree) rotateLeft(node *AvlNode) *AvlNode {
	rightNode := node.Right
	node.Right = rightNode.Left
	rightNode.Left = node
	// updating the height and balance factor for the rotated
	// nodes.
	avl.update(node)
	avl.update(rightNode)
	return rightNode
}

// Search walks through the avl tree to look for
// the specified item.
//
// it returns nil if the item does not exist.
func (avl *AvlTree) Search(item float64) *AvlNode {
	return avl.find(avl.root, item)
}

// find is a helper method to recursively find an item in
// the avl tree.
func (avl *AvlTree) find(node *AvlNode, item float64) *AvlNode {
	if node == nil {
		return nil
	}
	if item > node.Data {
		return avl.find(node.Right, item)
	}
	if item < node.Data {
		return avl.find(node.Left, item)
	}
	return node
}

// Size returns the size of the avl tree.
func (avl *AvlTree) Size() int {
	return avl.size
}

// Remove removes an item from the avl tree.
func (avl *AvlTree) Remove(item float64) bool {
	if avl.find(avl.root, item) == nil {
		return false
	}
	avl.root = avl.removeItem(avl.root, item)
	avl.size--
	return true
}

// removeItem is a helper function to remove a node from the avl
// tree.
func (avl *AvlTree) removeItem(node *AvlNode, item float64) *AvlNode {
	if item > node.Data {
		node.Right = avl.removeItem(node.Right, item)
	}
	if item < node.Data {
		node.Left = avl.removeItem(node.Left, item)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		if node.Left.height > node.Right.height {
			successor := avl.findMaxNode(node.Left)
			node.Data = successor.Data
			node.Left = avl.removeItem(node.Left, successor.Data)
		} else {
			successor := avl.findMinNode(node.Right)
			node.Data = successor.Data
			node.Right = avl.removeItem(node.Right, successor.Data)
		}
	}
	avl.update(node)
	return avl.balance(node)
}

// findMinNode is a helper function to find min child node by
// digging left in a subtree.
func (avl *AvlTree) findMinNode(node *AvlNode) *AvlNode {
	if node.Left == nil {
		return node
	}
	return avl.findMinNode(node.Left)
}

// findMaxNode is a helper function to find max child node in
// a subtree by digging right.
func (avl *AvlTree) findMaxNode(node *AvlNode) *AvlNode {
	if node.Right == nil {
		return node
	}
	return avl.findMaxNode(node.Right)
}

// GetRoot returns the root of the avl tree.
func (avl *AvlTree) GetRoot() *AvlNode {
	return avl.root
}

// PreOrderTraversal runs a pre-order traversal on the avl tree
// and execute the callback function f for each iteration.
//
// pre-order traversal follows the order: <root>-<left>-<right>
func (avl *AvlTree) PreOrderTraversal(f func(node *AvlNode)) {
	avl.runPreOrderTraversal(avl.root, f)
}

func (avl *AvlTree) runPreOrderTraversal(node *AvlNode, f func(node *AvlNode)) {
	if node == nil {
		return
	}
	f(node)
	avl.runPreOrderTraversal(node.Left, f)
	avl.runPreOrderTraversal(node.Right, f)
}

// InOrderTraversal runs an in-order traversal on the avl tree and
//  execute the callback function f for each iteration.
//
// in-order traversal follows the order: <left>-<root>-<right>
func (avl *AvlTree) InOrderTraversal(f func(node *AvlNode)) {
	avl.runInOrderTraversal(avl.root, f)
}

func (avl *AvlTree) runInOrderTraversal(node *AvlNode, f func(node *AvlNode)) {
	if node == nil {
		return
	}
	avl.runInOrderTraversal(node.Left, f)
	f(node)
	avl.runInOrderTraversal(node.Right, f)
}

// PostOrderTraversal runs a post-order traversal on the avl tree and
// execute the callback function f for each iteration.
//
// post-order traversal follows the order: <left>-<right>-<root>
func (avl *AvlTree) PostOrderTraversal(f func(node *AvlNode)) {
	avl.runPostOrderTraversal(avl.root, f)
}

func (avl *AvlTree) runPostOrderTraversal(node *AvlNode, f func(node *AvlNode)) {
	if node == nil {
		return
	}
	avl.runPostOrderTraversal(node.Left, f)
	avl.runPostOrderTraversal(node.Right, f)
	f(node)
}

// LevelOrderTraversal runs a level-order traversal on the avl
// tree and execute the callback function f for each iteration.
func (avl *AvlTree) LevelOrderTraversal(f func(node *AvlNode)) {
	avl.runLevelOrderTraversal(avl.root, f)
}

func (avl *AvlTree) runLevelOrderTraversal(node *AvlNode, f func(node *AvlNode)) {
	if node == nil {
		return
	}
	queue := NewQueue()
	queue.Enqueue(node)
	for !queue.IsEmpty() {
		currentItem, _ := queue.Dequeue()
		currentNode := currentItem.(*AvlNode)
		f(currentNode)

		if currentNode.Left != nil {
			queue.Enqueue(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.Enqueue(currentNode.Right)
		}
	}
}
