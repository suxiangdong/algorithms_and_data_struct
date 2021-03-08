package dqueue

import "errors"

type DQueue interface {
	FirstPush(elem)
	FirstPop() (elem, error)
	LastPush(elem)
	LastPop() (elem, error)
	Len() int
}

type elem interface {}

// slice 实现
type sliceDQueue struct {
	arr []elem
	size int
}

func New() *sliceDQueue {
	return &sliceDQueue{
		arr: make([]elem, 0),
		size: 0,
	}
}

func (d *sliceDQueue) FirstPush(e elem) {
	arr := make([]elem, d.Len() + 1)
	copy(arr[1:], d.arr)
	arr[0] = e
	d.arr = arr
	d.size++
}

func (d *sliceDQueue) FirstPop() (elem, error) {
	if d.Len() > 0 {
		elem := d.arr[0]
		d.arr = d.arr[1:]
		d.size--
		return elem, nil
	}

	return nil, errors.New("empty queue")
}

func (d *sliceDQueue) LastPush(e elem) {
	d.arr = append(d.arr, e)
	d.size++
}

func (d *sliceDQueue) LastPop() (elem, error) {
	if d.Len() > 0 {
		d.size--
		elem := d.arr[d.Len()]
		d.arr = d.arr[0: d.Len()]
		return elem, nil
	}

	return nil, errors.New("empty queue")
}

func (d *sliceDQueue) Len() int {
	return d.size
}