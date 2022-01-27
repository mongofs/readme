package stack

import (
	"errors"
	"readme/arithmetic/linkedlist"
)

type LinkStack struct {
	content   linkedlist.LinkedList
	n , limit int
}

func (l *LinkStack) Push(i interface{}) error {
	if l.limit ==l.n {return errors.New("capacity is nil ")}
	l.content.PushBack(&linkedlist.Node{Data: i,})
	l.n ++
	return nil
}

func (l *LinkStack) Pop() interface{} {
	if l.n == 0 {return nil }
	result := l.content.RemoveLast()
	l.n --
	return result
}

