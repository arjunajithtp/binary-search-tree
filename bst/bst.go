package bst

// constants to determine the traversal order
const (
	Preorder  = 1
	Postorder = 2
	Inorder   = 3
)

type node struct {
	data       int
	leftChild  *node
	rightChild *node
}

// BST is our Binary Search Tree
type BST struct {
	root *node
}

// Insert method will insert the data into the BST
func (b *BST) Insert(data ...int) {
	for _, new := range data {
		newNode := &node{
			data: new,
		}
		if b.root == nil { // this is the root node
			b.root = newNode
			continue
		}
		b.root.insert(newNode)
	}
}

func (n *node) insert(newNode *node) {
	if newNode.data <= n.data { // going left
		if n.leftChild == nil {
			n.leftChild = newNode
			return
		}
		n.leftChild.insert(newNode)
		return
	}
	// going right
	if n.rightChild == nil {
		n.rightChild = newNode
		return
	}
	n.rightChild.insert(newNode)
}

// CollectData will return a slice with data collected in selected traversal method
func (b *BST) CollectData(travesalMethod int) []int {
	var result []int

	switch travesalMethod {
	case Preorder:
		result = b.root.preorder(result)
	case Postorder:
		result = b.root.postorder(result)
	case Inorder:
		result = b.root.inorder(result)
	}

	return result
}

func (n *node) preorder(result []int) []int {
	if n == nil {
		return result
	}

	result = append(result, n.data)        // root
	result = n.leftChild.preorder(result)  // left traversal
	result = n.rightChild.preorder(result) // right traversal

	return result
}

func (n *node) postorder(result []int) []int {
	if n == nil {
		return result
	}

	result = n.leftChild.postorder(result)  // left traversal
	result = n.rightChild.postorder(result) // right traversal
	result = append(result, n.data)         // root

	return result
}

func (n *node) inorder(result []int) []int {
	if n == nil {
		return result
	}

	result = n.leftChild.inorder(result)  // left traversal
	result = append(result, n.data)       // root
	result = n.rightChild.inorder(result) // right traversal

	return result
}
