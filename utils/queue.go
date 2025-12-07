package utils

type Queue[T any] struct {
	head *listNode[T]
	tail *listNode[T]
}

type listNode[T any] struct {
	val  T
	next *listNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Enqueue(val T) {
	if q.IsEmpty() {
		n := &listNode[T]{val: val}
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = &listNode[T]{val: val}
	q.tail = q.tail.next
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var empty T
		return empty, false
	}

	v := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	return v.val, true
}
