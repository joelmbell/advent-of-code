package datastructures

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func (s Set[T]) Intersection(input Set[T]) Set[T] {

	largest := s
	smallest := input
	if len(input) > len(s) {
		largest = input
		smallest = s
	}

	output := make(Set[T])
	for a := range largest {
		if _, ok := smallest[a]; ok {
			output[a] = struct{}{}
		}
	}

	return output
}
