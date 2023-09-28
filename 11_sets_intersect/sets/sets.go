package sets

func NewSet[T comparable](elements ...T) *Set[T] {
	newSet := Set[T]{
		m: make(map[T]struct{}, len(elements)),
	}

	for _, element := range elements {
		newSet.m[element] = struct{}{}
	}

	return &newSet
}

// Множество будем реализовывать с помощью мапы
type Set[T comparable] struct {
	m map[T]struct{}
}

// Возвращает мощность множества
// (простыми словами - количество элементов)
func (s *Set[T]) cardinality() int {
	return len(s.m)
}

// Возвращает слайс из всех элементов множества
// Порядок элементовне гарантирован
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.m))

	for element := range s.m {
		elements = append(elements, element)
	}

	return elements
}

// возвращает новое множество -- результат пересечения двух множеств
func (s1 *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	newSet := Set[T]{
		m: map[T]struct{}{},
	}

	// небольшая оптимизация
	if s1.cardinality() > s2.cardinality() {
		for elem := range s2.m {
			if _, ok := s1.m[elem]; ok {
				newSet.m[elem] = struct{}{}
			}
		}

		return &newSet
	}

	for elem := range s1.m {
		if _, ok := s2.m[elem]; ok {
			newSet.m[elem] = struct{}{}
		}
	}

	return &newSet
}
