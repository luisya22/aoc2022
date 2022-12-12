package datastruct

type Queue[T comparable] struct {
	Data []T
}

func (q *Queue[T]) Dequeue() *T {
	if len(q.Data) == 0 {
		return nil
	}

	item := q.Data[0]
	q.Data = q.Data[1:]

	return &item
}

func (q *Queue[T]) Queue(item T) {
	q.Data = append(q.Data, item)
}
