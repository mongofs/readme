// 由于golang 中的slice 切片类型已经自己支持了扩容功能，所以在golang中实现栈的迁移还是比较简单的
// 最好不要基于本结构做工业用途
package stack

import "errors"

type ArrayStack struct {
	content []interface{} // 具体存放的地方
	n       int           // 元素的个数
	limit   int
}

func NewArrayStack(i int) *ArrayStack {
	return &ArrayStack{
		content: nil,
		n:       0,
		limit:   i,
	}
}

func (s *ArrayStack) Push(data interface{}) error{
	if s.limit == s.n {
		return errors.New("capacity is nil ")
	}
	s.content = append(s.content,data)
	s.n++
	return nil
}


func (s *ArrayStack)Pop ()interface{}{
	if s.n == 0 {return nil}
	res := s.content[s.n-1]
	s.content = s.content [0:len(s.content)-1]
	s.n--
	return res
}