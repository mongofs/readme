package linklist

// 反转链表
func (l *list)Reserve(){
	// 反转链表
	var curr,prev   *Node = l.head,nil
	//  1->2->3->
	// <- 1<-2 <- 3
	for curr != nil {
		next := curr.next
		curr.next =prev
		prev = curr
		curr =next
	}
	l.head =prev
}


// 链表排序


//

