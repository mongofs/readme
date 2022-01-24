// 由于golang 中的slice 切片类型已经自己支持了扩容功能，所以在golang中实现栈的迁移还是比较简单的


type ArrayStack struct {
	content [] interface {} // 具体存放的地方
	n int 	// 元素的个数
}

func NewArrayStack ()*ArrayStack{
	
}

func (s *ArrayStack)Push (){


}