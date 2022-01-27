package queue

import "readme/arithmetic/linkedlist"

type ListQueue struct {
	list linkedlist.LinkedList
	size int
}

func (l *ListQueue) Enqueue(i interface{}) bool {
	l.list.PushBack(&linkedlist.Node{Data: i})
	l.size ++
	return true
}

func (l ListQueue) Dequeue() interface{} {
	node := l.list.RemoveHead()
	if node !=nil {return nil}
	l.size--
	return node.Data
}

func (l ListQueue) Size() int {
	return l.size
}


