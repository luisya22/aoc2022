package datastruct

type Queue[T comparable] struct {
	Data []T
}

func (q *Queue[T]) Dequeue() {
	q.Data = q.Data[1:]
}

func (q *Queue[T]) Queue(item T) {
	q.Data = append(q.Data, item)
}
