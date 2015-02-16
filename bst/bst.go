package bst

type Node struct {
	Key   int
	Value interface{}
	left  *Node
	right *Node
}

func (n *Node) LeftChild() *Node {
	return n.left
}

func (n *Node) RightChild() *Node {
	return n.right
}

func (n *Node) insert(node *Node) interface{} {
	if node.Key <= n.Key {
		if n.left == nil {
			n.left = node
			return node.Value
		} else {
			return n.left.insert(node)
		}
	} else {
		if n.right == nil {
			n.right = node
			return node.Value
		} else {
			return n.right.insert(node)
		}
	}
}

type BST struct {
	root *Node
	size int
}

func New() *BST { return new(BST).Init() }

func (b *BST) Init() *BST {
	b.root = nil
	b.size = 0
	return b
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) Root() *Node {
	if b.size == 0 {
		return nil
	}
	return b.root
}

func (b *BST) RootKey() int {
	if b.size == 0 {
		return -1
	}
	return b.root.Key
}

func (b *BST) RootValue() interface{} {
	if b.size == 0 {
		return nil
	}
	return b.root.Value
}

func (b *BST) Insert(key int, value interface{}) interface{} {
	b.size++
	return b.insert(key, value)
}

func (b *BST) insert(key int, value interface{}) interface{} {
	n := Node{Key: key, Value: value}
	if b.root == nil {
		b.root = &n
		return value
	}
	return b.root.insert(&n)
}
