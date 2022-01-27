package linkedlist

// 有序链表和那个之前设计的单链表节点不同，为了方便演示就不用list了
type sortNode struct {
	next *sortNode
	data int
}


// 合并两个有序链,返回较小的Node
func MergeSortNode(n1,n2 *sortNode)  *sortNode{
	if  n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	// 如果n1 的值大于n2 的值
	if n1.data<n2.data {
		n1.next = MergeSortNode(n1.next,n2)
		return n1
	}else {
		n2.next = MergeSortNode(n1,n2.next)
		return n2
	}
}


// 关于链表排序 逻辑
//

func midNode(head *sortNode) *sortNode{
	if head!= nil || head.next == nil{
		return head
	}
	var slow,fast = head,head.next

	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	mid := slow.next
	slow.next = nil
	return mid
}



//递归将链表分成两段，合并两条有序链表
func SortList(head *sortNode) *sortNode {
	if head == nil || head.next == nil {
		return head
	}
	//
	var temhead ,middle = head,midNode(head)
	temhead = SortList(temhead)
	middle = SortList(middle)
	return MergeSortNode(temhead, middle)
}
