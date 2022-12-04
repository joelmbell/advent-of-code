package datastructures

import "testing"

func TestOneNode(t *testing.T) {
	list := NewCircularLinkedList("Hello")
	head := list.Head
	if head.Value != "Hello" {
		t.Errorf("Invalid Head data")
	}

	if head != head.Previous {
		t.Errorf("single node lists previous should be the same as head")
	}

	if head != head.Next {
		t.Errorf("single node lists next should be the same as head")
	}

	if list.Len != 1 {
		t.Errorf("length should be 1")
	}
}

func TestHeadWithTwoNodes(t *testing.T) {
	list := NewCircularLinkedList("Hello", "Goodbye")

	if list.Head.Value != "Hello" {
		t.Errorf("invalid head data after adding second node")
	}
	if list.Head.Next.Value != "Goodbye" {
		t.Errorf("head next should reference the second node")
	}

	if list.Head.Previous.Value != "Goodbye" {
		t.Errorf("head previous should reference the second node")
	}

	if list.Len != 2 {
		t.Errorf("invalid length, should be 2")
	}
}

func TestTailWithTwoNodes(t *testing.T) {
	list := NewCircularLinkedList("Hello", "Goodbye")

	if list.Tail.Value != "Goodbye" {
		t.Errorf("tail is not pointing to the latest node")
	}
	if list.Tail.Next.Value != "Hello" {
		t.Errorf("tail next is not pointing to head")
	}

	if list.Tail.Previous.Value != "Hello" {
		t.Errorf("tail previous is not pointing to the first node")
	}
}

func Test3Nodes(t *testing.T) {
	list := NewCircularLinkedList("A", "B", "C")

	if list.Head.Next.Value != "B" {
		t.Errorf("head not pointing to second element")
	}

	if list.Head.Next.Next.Value != "C" {
		t.Errorf("head.Next.Next not pointing to 3rd element")
	}

	if list.Head.Next.Next.Next.Value != "A" {
		t.Errorf("head.Next.Next.Next.data not pointing to itself")
	}

	if list.Tail.Value != "C" {
		t.Errorf("tail not correct value")
	}

	if list.Tail.Previous.Value != "B" {
		t.Errorf("tail.Previous not pointing to second element")
	}

	if list.Tail.Previous.Previous.Value != "A" {
		t.Errorf("tail.Previous.Previous not pointing to 1st element")
	}

	if list.Tail.Previous.Previous.Previous.Value != "C" {
		t.Errorf("tail.Previous.Previous.Previous not pointing to itself")
	}

	if list.Len != 3 {
		t.Errorf("invalid length, should be 3")
	}
}

func TestGetNode(t *testing.T) {
	list := NewCircularLinkedList("A", "B", "C")
	node := list.Find("B")

	if node.Value != "B" {
		t.Errorf("found invalid node")
	}

	if node.Previous.Value != "A" {
		t.Errorf("node has invalid previous")
	}

	if node.Next.Value != "C" {
		t.Errorf("node has invalid next")
	}
}
