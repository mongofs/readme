package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_Push(t *testing.T) {
	l := New()
	l.Push(&Node{data: 1})
	l.Push(&Node{data: 2})
	l.Push(&Node{data: 3})
	fmt.Println(l.Range())

	// output : 1,2,3
}

func TestList_Remove(t *testing.T) {
	l := New()
	data1 :=&Node{data: 1}
	data2 :=&Node{data: 2}
	data3 :=&Node{data: 3}
	data4 :=&Node{data: 4}

	l.Push(data1)
	l.Push(data2)
	l.Push(data3)
	l.Push(data4)
	fmt.Println(l.Range())
	// output : 1,2,3,4

	l.Remove(data2)
	l.Remove(data3)
	fmt.Println(l.Range())
	// output : 1,4

	l.Remove(data1)
	fmt.Println(l.Range())
	// output : 4

	l.Remove(&Node{data: 10})
	fmt.Println(l.Range())
	// output : 4

	l.Remove(data4)
	fmt.Println(l.Range())
	// output : []
}

func TestList_RmOrderbyDesc(t *testing.T) {
	l := New()
	l.Push(&Node{data: 1})
	l.Push(&Node{data: 2})
	l.Push(&Node{data: 3})
	l.Push(&Node{data: 4})
	l.Push(&Node{data: 5})
	fmt.Println(l.Range())

	// output : 1,2,3

	l.RmOrderbyDesc(2)
	fmt.Println(l.Range())
	// output :1,2
}
