package datastruct

type Stack[T comparable] struct {
	Data []T
}

func (s *Stack[T]) Push(item T) {
	s.Data = append(s.Data, item)
}

func (s *Stack[T]) Pop() T {
	removedItem := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return removedItem
}

func (s *Stack[T]) Peek() *T {
	if len(s.Data) > 0 {
		return &s.Data[len(s.Data)-1]
	}

	return nil
}
