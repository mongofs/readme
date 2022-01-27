package queue



type ArrayQueue struct {
	content []interface{}
	head,tail,size int
}


func (a *ArrayQueue) Enqueue(i interface{}) bool{
	if a.tail == a.size  {
		// 队尾满了
		if a.head == 0 {
			// 对头也满了
			return false
		}

		for i:= a.head ;i<a.size;i++{
			a.content[i-a.head] = a.content[i]
		}
		a.head = 0
		a.tail = a.tail -a.head
	}
	a.content[a.tail] = i
	a.tail ++
	return true
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.head ==a.tail {return nil }
	res := a.content[a.head]
	a.content= a.content[a.head+1:]
	return res
}


func (a *ArrayQueue) Size() int {
	return a.size
}





