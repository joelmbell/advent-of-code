package datastructures

import "testing"

func TestSet_IntersectionWithValues(t *testing.T) {
	a := NewSet([]string{"A", "B", "C"})
	b := NewSet([]string{"A", "C"})

	result := a.Intersection(b)
	if len(result) != 2 {
		t.Errorf("should have 2 elements has %v", len(result))
	}

	result2 := b.Intersection(a)
	if len(result2) != 2 {
		t.Errorf("should have 2 elements has %v", len(result2))
	}
}

func TestSet_IntersectionWithoutValues(t *testing.T) {
	a := NewSet([]string{})
	b := NewSet([]string{"A", "B", "C"})

	result := a.Intersection(b)
	if len(result) != 0 {
		t.Errorf("result should have 0 elements, has %v", len(result))
	}
}
