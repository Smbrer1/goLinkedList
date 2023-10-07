package main

import "fmt"

type LinkedList[T any] struct {
	head node[T]
}

type node[T any] struct {
	value T
	ptr   *node[T]
}

func (ll *LinkedList[T]) push(elem T) {
	t := node[T]{elem, &ll.head}
	fmt.Println(t.ptr)
	ll.head = t
}

func newList[T any]() LinkedList[T] {
	return LinkedList[T]{
		head: node[T]{},
	}
}

func main() {
	ll := newList[int]()

	fmt.Println(ll)
	ll.push(2)
	fmt.Println(ll)
	fmt.Println(ll.head.ptr)
}
