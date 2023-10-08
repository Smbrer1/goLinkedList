package golinkedlist

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyList(t *testing.T) {
	l := NewEmptyList[int]()

	assert.Equal(
		t,
		LinkedList[int]{head: nil},
		l,
		"Function NewList does not return list with empty node",
	)
}

func TestNewList(t *testing.T) {
	l := NewList[int](1)
	val := 1
	assert.Equalf(
		t,
		LinkedList[int]{head: &node[int]{value: val}},
		l,
		"Function NewEmptyList does not return list with first node value: %v",
		val,
	)
	assert.Equalf(t, val, l.head.value, "Value of first node does not equal %v", val)
}

func TestPush(t *testing.T) {
	l := LinkedList[int]{}
	val1, val2, val3 := 1, 2, 3
	l.Push(val1)
	l.Push(val2)
	l.Push(val3)
	assert.Equalf(t, val3, l.head.value, "Value of first node does not equal %v", val3)
	assert.Equalf(t, val2, l.head.next.value, "Value of second node does not equal %v", val2)
	assert.Equalf(t, val1, l.head.next.next.value, "Value of third node does not equal %v", val1)
}

func TestPop(t *testing.T) {
	val1, val2 := 1, 2
	l := LinkedList[int]{head: &node[int]{value: val1, next: &node[int]{value: val2}}}
	v, err := l.Pop()
	assert.Equalf(t, val1, v, "Value of popped element is not equals %v", val1)
	assert.Equal(
		t,
		LinkedList[int]{head: &node[int]{value: val2}},
		l,
		"Popped list does no join 2nd node",
	)
	v, err = l.Pop()
	assert.Equalf(t, val2, v, "Value of popped element is not equals %v", val2)
	assert.Equal(t, LinkedList[int]{}, l, "Popped list does not join nil node")
	v, err = l.Pop()
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("LinkedList is empty"), err, "LinkedList is not empty")
	}
}

func TestIter(t *testing.T) {
	val1, val2, val3 := 1, 2, 3
	l_empty := LinkedList[int]{}
	l := LinkedList[int]{
		head: &node[int]{
			value: val1, next: &node[int]{
				value: val2, next: &node[int]{
					value: val3,
				},
			},
		},
	}

	i_nil, err := l_empty.Iter()
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("LinkedList is empty"), err, "")
		assert.Nil(t, i_nil, "Iter is not nil")
	}

	i, err := l.Iter()
	if err != nil {
		panic(err)
	}

	n1 := i()
	n2 := i()
	n3 := i()
	assert.Equalf(t, val1, n1.value, "Value in first node is not equal to %v", val1)
	assert.Equalf(t, val2, n2.value, "Value in second node is not equal to %v", val2)
	assert.Equalf(t, val3, n3.value, "Value in third node is not equal to %v", val3)

	n_nil := i()
	assert.Nil(t, n_nil, "Iter does not return nil if list is empty")
}
