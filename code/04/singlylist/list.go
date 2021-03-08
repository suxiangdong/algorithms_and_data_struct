package singlylist

import (
	"errors"
)

type List interface {
	Append(Elem)
	Insert(int, Elem) bool
	Get(int) (Elem, error)
	Delete(int) bool
	Len() int
	ToArray() []Elem
}

type Elem interface{}

type node struct {
	data Elem
	next *node
}

type linkedList struct {
	head *node
	tail *node
	length int
}

func New() *linkedList {
	return &linkedList{
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (list *linkedList) Len() int {
	return list.length
}

func (list *linkedList) ToArray() []Elem {
	arr := make([]Elem, list.length)
	node := list.head
	i := 0
	for {
		arr[i] = node.data
		if node.next == nil {
			break
		}
		i++
		node = node.next
	}
	return arr
}

// 尾部添加
func (list *linkedList) Append(elem Elem) {
	node := &node{data: elem, next: nil}
	if list.length == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
	list.length++
}

func (list *linkedList) Insert(i int, elem Elem) bool {
	if i > list.length + 1 {
		return false
	}
	node := &node{data: elem, next: nil}

	if i == 1 {
		node.next = list.head
		list.head = node
	} else {
		prevNode := list.head
		for j := 1; j < i - 1; j++ {
			prevNode = prevNode.next
		}

		node.next = prevNode.next
		prevNode.next = node
	}

	if i == list.length + 1 {
		list.tail = node
	}
	list.length++

	return true
}

func (list *linkedList) Delete(i int) bool {
	if i > list.length {
		return false
	}

	if i == 1 {
		list.head = list.head.next
	}

	prevNode := list.head
	for j := 1; j < i - 1; j++ {
		prevNode = prevNode.next
	}
	beDeletedNode := prevNode.next
	nextNode := beDeletedNode.next
	prevNode.next = nextNode
	list.length--

	if nextNode == nil {
		list.tail = prevNode
	}
	return true
}

func (list *linkedList) Get(i int) (Elem, error) {
	if i > list.length {
		return 0, errors.New("out bound")
	}

	node := list.head
	for j := 1; j < i; j++ {
		node = node.next
	}
	return node.data, nil
}