package datastructures

type Node[V comparable] struct {
	Value    V
	Next     *Node[V]
	Previous *Node[V]
}

type CircularLinkedList[T comparable] struct {
	Len  int
	Head *Node[T]
	Tail *Node[T]
}

func NewCircularLinkedList[T comparable](values ...T) *CircularLinkedList[T] {
	list := &CircularLinkedList[T]{}
	for _, value := range values {
		list.Add(value)
	}
	return list
}

func (c *CircularLinkedList[V]) Add(data V) {
	node := &Node[V]{
		Value: data,
	}

	if c.Head == nil {
		c.Head = node
		c.Tail = node
		c.Head.Next = node
		c.Head.Previous = node
	}

	c.Tail.Next = node
	node.Next = c.Head
	node.Previous = c.Tail
	c.Tail = node
	c.Head.Previous = node

	c.Len += 1
}

func (c *CircularLinkedList[T]) Find(value T) *Node[T] {
	i := 0
	current := c.Head
	for i < c.Len {
		if current.Value == value {
			return current
		}
		current = current.Next
		i++
	}
	return nil
}
