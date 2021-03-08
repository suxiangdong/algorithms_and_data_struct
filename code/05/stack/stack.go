package stack

import "errors"
import "github.com/suxiangdong/algorithms_and_data_struct/code/04/singlylist"

type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	Len() int
}

// 数组（slice）实现
type sliceStack struct {
	arr []interface{}
	len int
}

// 链表实现（linkedList）实现【借助第四节实现的链表】
type linkedListStack struct {
	list singlylist.List
}

func NewSliceStack() *sliceStack {
	return &sliceStack{
		arr: make([]interface{}, 0),
		len: 0,
	}
}

func (stack *sliceStack) Push(elem interface{}) {
	stack.arr = append(stack.arr, elem)
	stack.len++
}

func (stack *sliceStack) Pop() (interface{}, error) {
	if stack.Len() > 0 {
		stack.len--
		elem := stack.arr[stack.Len()]
		stack.arr = stack.arr[:stack.len]
		return elem, nil
	}
	return nil, errors.New("no element")
}

func (stack *sliceStack) Len() int {
	return stack.len
}


func NewLinkedListStack() *linkedListStack {
	return &linkedListStack{
		list: singlylist.New(),
	}
}

func (stack *linkedListStack) Push(elem interface{}) {
	stack.list.Append(elem)
}

func (stack *linkedListStack) Pop() (interface{}, error) {
	if length := stack.Len(); length > 0 {
		elem, _ := stack.list.Get(length)
		stack.list.Delete(length)
		return elem, nil
	}
	return nil, errors.New("no element")
}

func (stack *linkedListStack) Len() int {
	return stack.list.Len()
}