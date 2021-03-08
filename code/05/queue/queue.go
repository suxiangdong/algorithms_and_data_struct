package queue

import "errors"

type Queue interface {
	Push(elem)
	Pop() (elem, error)
	Len() int
}

type elem interface{}

// slice 实现
type sliceQueue struct {
	arr []elem
	size int
}

func New() *sliceQueue {
	return &sliceQueue{
		arr: make([]elem, 0),
		size: 0,
	}
}

func (q *sliceQueue) Push(e elem) {
	q.arr = append(q.arr, e)
	q.size++
}

func (q *sliceQueue) Pop() (elem, error) {
	if q.Len() > 0 {
		elem := q.arr[0]
		q.arr = q.arr[1:]
		q.size--
		return elem, nil
	}

	return nil, errors.New("empty queue")
}

func (q *sliceQueue) Len() int {
	return q.size
}