package linklist

import (
	"fmt"
	"testing"
)

func TestMergeSortNode(t *testing.T) {
	sort1 := &sortNode{data: 1, next: &sortNode{next: &sortNode{next: nil, data: 5}, data: 2,}}
	sort2 := &sortNode{data: 2, next: &sortNode{next: &sortNode{next: nil, data: 6}, data: 3,}}
	result := MergeSortNode(sort1,sort2)
	fmt.Println(result)
}


func TestSortList(t *testing.T) {
	n5 := &sortNode{next: nil, data:5,}
	n4 := &sortNode{next: n5, data: 4,}
	n3 := &sortNode{next: n4, data: 3,}
	n2 := &sortNode{next: n3, data: 2,}
	n1 := &sortNode{next: n2, data: 1,}

	head := &sortNode{
		next: n1,
		data: 0,
	}
	newhead := SortList(head)
	fmt.Println(newhead)
}