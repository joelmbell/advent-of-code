package datastructures

type TreeNodeValue interface {
	Id() string
}

type TreeNode[T TreeNodeValue] struct {
	Value    T
	Children []*TreeNode[T]
	Parent   *TreeNode[T]
}

type Tree[T TreeNodeValue] struct {
	Root *TreeNode[T]
}

func NewTree[T TreeNodeValue](value T) *Tree[T] {
	return &Tree[T]{
		Root: NewTreeNode[T](value),
	}
}

func NewTreeNode[T TreeNodeValue](value T) *TreeNode[T] {
	return &TreeNode[T]{
		Value:    value,
		Children: make([]*TreeNode[T], 0),
	}
}

func (n *TreeNode[T]) Add(value T) *TreeNode[T] {
	newNode := NewTreeNode(value)
	newNode.Parent = n
	n.Children = append(n.Children, newNode)
	return newNode
}

func (n *TreeNode[T]) FindChild(value T) *TreeNode[T] {
	for _, child := range n.Children {
		if child.Value.Id() == value.Id() {
			return child
		}
	}
	return nil
}

func (n *TreeNode[T]) Iterate(iterator func(*TreeNode[T], int), level int) {
	iterator(n, level)

	if len(n.Children) <= 0 {
		return
	}

	for _, child := range n.Children {
		nextLevel := level + 1
		child.Iterate(iterator, nextLevel)
	}
}
