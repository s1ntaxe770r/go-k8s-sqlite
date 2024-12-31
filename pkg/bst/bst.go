package bst

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(node *Node, val int) *Node {
	// If root node is nill, return a new node
	if bst.Root == nil {
		bst.Root = &Node{val, nil, nil}
		return bst.Root
	}
	// node might be empty on first iteration so check
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.Insert(node.left, val)
	}
	if val > node.data {
		node.right = bst.Insert(node.right, val)
	}
	return node
}

// Print each level of the bst
func (bst *BinarySearchTree) LevelOrder() []int {
	var result []int

	queue := []*Node{bst.Root}

	for len(queue) > 0 {
		current := queue[0]

		queue = queue[1:]

		result = append(result, current.data)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		// If the right child exists, add it to the queue
		if current.right != nil {
			queue = append(queue, current.right)
		}

	}
	return result
}
