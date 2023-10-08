package golinkedlist

type node[T any] struct {
	value T
	next  *node[T]
}
