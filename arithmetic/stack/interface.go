package stack

type Stack interface {
	Push(interface{})error
	Pop()interface{}
}