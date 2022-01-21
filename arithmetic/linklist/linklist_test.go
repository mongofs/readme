package linklist

import (
	"fmt"
	"testing"
)

func TestList_Reserve(t *testing.T) {
	l := New()
	l.Push(&Node{data: 1})
	l.Push(&Node{data: 2})
	l.Push(&Node{data: 3})
	l.Push(&Node{data: 4})
	l.Push(&Node{data: 5})
	fmt.Println(l.Range())
	// output 1,2,3,4,5

	l.Reserve()
	fmt.Println(l.Range())
	// output 5,4,3,2,1
}
