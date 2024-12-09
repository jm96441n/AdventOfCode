package utils

import (
	"cmp"
)

type PriorityQueue[T cmp.Ordered] struct {
	Size int
	heap []T
}

func ToPriorityQueue[T cmp.Ordered](arr []T) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{heap: make([]T, 0, len(arr))}
	for _, v := range arr {
		pq.Push(v)
	}
	return pq
}

func (pq *PriorityQueue[T]) Push(value T) {
	pq.heap = append(pq.heap, value)
	pq.Size += 1

	idx := len(pq.heap) - 1
	for idx > 0 {
		ele := pq.heap[idx]
		parentIdx := (idx - 1) / 2
		parent := pq.heap[parentIdx]

		if ele > parent {
			break
		}
		pq.heap[idx], pq.heap[parentIdx] = parent, ele
		idx = parentIdx
	}
}

func (pq *PriorityQueue[T]) Pop() T {
	val := pq.heap[0]
	pq.Size -= 1
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[0 : len(pq.heap)-1]

	idx := 0
	min := idx

	for idx < len(pq.heap) {
		left := (2 * idx) + 1
		right := left + 1

		if (left < len(pq.heap) && pq.heap[left] < pq.heap[min]) || (right < len(pq.heap) && pq.heap[right] < pq.heap[min]) {
			if right < len(pq.heap) {
				min = right
				if pq.heap[left] < pq.heap[right] {
					min = left
				}
			} else {
				min = left
			}
		}
		if idx == min {
			break
		}
		pq.heap[min], pq.heap[idx] = pq.heap[idx], pq.heap[min]
		idx = min
	}
	return val
}
