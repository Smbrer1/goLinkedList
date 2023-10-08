package main

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

func (ll *LinkedList[T]) Push(elem T) {
	list := &node[T]{elem, ll.head}
	ll.head = list
}

func (ll *LinkedList[T]) Pop() (T, error) {
	if ll.head == nil {
		return *new(T), errors.New("LinkedList is empty")
	} else {
		n := ll.head
		ll.head = ll.head.next
		return n.value, nil
	}
}

func NewList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: &node[T]{},
	}
}

func (ll LinkedList[T]) iter() (func() node[T], error) {
	if ll.head != nil {
		n := ll.head
		return func() node[T] {
			t_node := n
			if n.next != nil {
				n = n.next
			}
			return *t_node
		}, nil
	}

	return nil, errors.New("LinkedList is empty")
}

func main() {
	ll := NewList[int]()
	printList[int](ll.head)
	ll.Push(2)
	printList[int](ll.head)
	printList[int](ll.head.next)
	ll.Push(3)
	printList[int](ll.head.next.next)
	p, _ := ll.Pop()
	printList[int](ll.head)
	fmt.Println(p)
	ll.Push(5)
	iter, _ := ll.iter()
	v := iter()
	printList[int](&v)
	v = iter()
	printList[int](&v)
	v = iter()
	iter()
	iter()
	iter()
	printList[int](&v)
	printList[int](ll.head)
}

func printList[T any](n *node[T]) {
	fmt.Printf("ptr: %p value: %v\n", n.next, n.value)
}
