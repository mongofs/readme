package stack

import (
	"fmt"
	"log"
	"testing"
)

var stack *ArrayStack = NewArrayStack(10)


func TestNewArrayStack(t *testing.T) {
	var tests = []struct {
		give interface{}
	}{
		{give: 1},
		{give: 2},
		{give: 3},
		{give: 4},
	}
	for _,v := range tests {
		if err := stack.Push(v.give);err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	// output 4,3,2,1
}


