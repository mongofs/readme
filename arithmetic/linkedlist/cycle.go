package linkedlist


// 检测是否有环 ,通过hash表判断
func DetectCycle(head *Node)*Node{
	mp := map[*Node]struct{}{}
	for head !=nil {
		if _,ok := mp[head]; ok  {
			return head
		}
		head = head.next
	}
	return nil
}

// 双指针检测是否存在环
// 检测原理 ：使用快慢指针，每次走的步数，慢指针始终都是快指针的1/2,如果存在环，那么慢指针和快指针始终会相遇
// 但是由于 慢指针 * 2  = 快指针的步数， 结合链表长度
func DetectCycle2(head *Node)*Node{
	if head == nil {return  nil}
	var slow ,fast  = head,head
	for fast !=nil {
		slow := slow.next
		if fast .next == nil {
			return nil
		}
		if fast == slow {
			p:= head
			for p !=slow {
				p= p.next
				slow = slow.next
			}
			return p
		}
	}
	return nil
}