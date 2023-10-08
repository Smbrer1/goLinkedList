package golinkedlist

import "errors"

type LinkedList[T any] struct {
	head *node[T]
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

func NewList[T any](elem T) LinkedList[T] {
	return LinkedList[T]{
		head: &node[T]{value: elem},
	}
}

func NewEmptyList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
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
