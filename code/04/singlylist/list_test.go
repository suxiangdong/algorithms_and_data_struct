package singlylist

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	observed := []Elem{1,2,3,4,5}
	list := New()
	for _, i := range observed {
		list.Append(i)
	}
	if list.tail.data != 5 {
		t.Fatal("append head failed")
	}
	if list.head.data != 1 {
		t.Fatal("append head failed")
	}
	if list.length != len(observed) {
		t.Fatal("append length failed")
	}
	if !reflect.DeepEqual(observed, list.ToArray()) {
		t.Fatal("append failed")
	}
}

func TestLinkedList_Insert(t *testing.T) {
	observed := []Elem{1,2,3,4,5}
	list := New()
	for k, i := range observed {
		list.Insert(k+1, i)
	}
	if list.tail.data != 5 {
		t.Fatal("insert head failed")
	}
	if list.head.data != 1 {
		t.Fatal("insert head failed")
	}
	if list.length != len(observed) {
		t.Fatal("insert length failed")
	}
	if !reflect.DeepEqual(observed, list.ToArray()) {
		t.Fatal("insert failed")
	}
}

func TestLinkedList_Delete(t *testing.T) {
	observed := []Elem{1,2,3,4,5}
	list := New()
	for k, i := range observed {
		list.Insert(k+1, i)
	}

	if ok := list.Delete(len(observed) + 1); ok {
		t.Fatal("delete out bound")
	}

	expected := []Elem{1,2,3,5}
	list.Delete(4)
	if !reflect.DeepEqual(expected, list.ToArray()) {
		t.Fatal("delete failed")
	}
}

func TestLinkedList_Get(t *testing.T) {
	observed := []Elem{1,2,3,4,5}
	list := New()
	for k, i := range observed {
		list.Insert(k+1, i)
	}

	if _,ok := list.Get(len(observed) + 1); ok == nil {
		t.Fatal("get out bound")
	}

	for k, elem := range observed {
		val, _ := list.Get(k+1)
		if elem != val {
			t.Fatal("get failed")
		}
	}
}