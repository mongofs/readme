package single

// 单链表
type List struct {
	head *Node
	length int
}

func NewList() *List {
	return  &List{
		head:   &Node{Next: nil},
		length: 0,
	}
}

// 将node 移动到Head头
func (l *List) MoveToFront(node *Node) {
	if node  == nil || node.list !=l {return }
	if l.head == nil { l.head =node ; return }
	node.Next =l.head
	l.head =node
}

// 将node 移动到尾部
func (l *List) MoveToBack(node *Node) {
	if node  == nil || node.list !=l {return }
	for tem:=l.head;tem!=nil;{
		if tem.Next == nil {
			tem.Next = node
			break
		}
		tem =tem.Next
	}
}

// 将node 移动到target 的前面
func (l *List) MoveBefore(node, target *Node) {
	if node  == nil || node.list !=l  || target == nil || target.list != l  {return }
	var tem = l.head
	for tem != nil {
		if tem.Next == target {
			break
		}
		tem =tem.Next
	}
	tem.Next = node
	node.Next =target
}

// 将node 移动到target 之前
func (l *List) MoveAfter(node, target *Node) {
	if node  == nil || node.list !=l  || target == nil || target.list != l  {return }
	tem := target.Next
	target.Next = node
	node.Next =tem
}

// 将node 插入到head头中 ,这里其实最好的处理办法需要判断一下node
// 是否在list中
func (l *List) PushFront(node *Node) {
	if node  == nil {return }
	if l.head == nil {
		l.head =node
		node.list =l
		return
	}
	node.Next =l.head
	node.list =l
	l.head =node
}

// 推送放入链表尾部
func (l *List) PushBack(node *Node) {
	if node  == nil {return }
	if l.head == nil {
		l.head =node
		node.list =l
		return
	}
	var tem =l.head
	for tem != nil {
		tem =tem.Next
	}
	tem.Next =node
	node.list =l
}

// 删除Node
func (l *List) RemoveNode(node *Node) {
	if node  == nil || node.list != l {return }
	var tem = l.head
	for tem!= nil {
		if tem.Next == node{
			tem.Next= node.Next
			break
		}
		tem=tem.Next
	}
}

// 删除倒数第k个Node
func (l *List) RemoveLastKthNode(lastKth int) *Node {
	if lastKth < 0 {return nil}
	 fast ,slow := l.head,l.head
	for i := 0;i<lastKth ;i++ {
		fast = fast.Next
	}
	var res  *Node
	for fast !=nil {
		if fast.Next == nil {
			res =slow.Next
			slow.Next =slow.Next.Next
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return res
}

// 删除尾结点
func (l *List) RemoveLast() *Node {
	if l.head == nil  {return nil}
	var tem = l.head
	for tem.Next !=nil {
		tem =tem.Next
	}
	res := tem.Next
	tem.Next = nil
	return res
}

// 删除头结点
func (l *List) RemoveHead() *Node {
	if l.head == nil  {return nil}
	tem :=l.head
	l.head = l.head.Next
	return tem
}

// 将链表进行排序 1-》2-》3-》 =》 3-》2-》1
func (l *List) Reverse() {
	if l.head == nil {return}
	var pre,curr  *Node = nil ,l.head
	for curr  !=nil {
		next :=curr.Next
		curr.Next =pre
		pre= curr
		curr =next
	}
}



