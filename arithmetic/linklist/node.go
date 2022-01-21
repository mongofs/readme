// node类的操作，node类操作并不是单纯的说是node结构体的操作方法，而是针对list上的node
// 进行操作。 由于一般工业使用都是使用的是双向链表，但是面试更多的是单链表的操作，所以这里
// 为了测试理论上的单链表变动，我在这个使用案例这里使用的是单向链表进行演示，这种链表没有实际应用
// 价值，如果为了工业化应用可以使用golang 的container包下的双向链表进行使用，我后期也会专门
// 编写一个stl 包进行工业化场景应用,所以在这个案例展示过程中部分方法我会选择性的省略掉，比如pushfront
// 等操作相似都将被去掉
package linklist

type Node struct {
	next *Node
	data interface{}
}

type list struct {
	head *Node // 指向链表的头结点
}

func New()*list {
	return &list{
		head: nil,
	}
}

// 节点类的操作：存在如下几种方式
// push
// remove
// move

// basic 操作
func (l*list) add(node *Node){
	if l.head == nil { l.head =node ;return }
	tem := l.head
	for tem.next != nil{
		tem = tem.next
	}
	tem.next = node
}

// 由于是单链表，删除头结点的时候和循环相比比较特殊，防止空指针，所以特殊判断
// 链表删除的时候需要特殊考虑的几个点
// 当链表没有节点
// 当链表节点为1
// 当链表节点为2
// 被删除值没有存在于

// O(n) 操作
func (l*list) rm(node *Node){
	// 节点数量为一
	if l.head.next ==nil  && node == l.head { l.head= nil ;return }
	// 节点数量为二
	if node == l.head { l.head =l.head.next ;return}
	tem := l.head
	for tem.next != nil  {
		if tem.next == node {
			// 判断node 节点是不是在尾结点上
			if tem.next.next == nil {
				tem.next = nil
			}else{
				tem.next=tem.next.next
			}
		}
		tem = tem.next
	}
}

// 将节点插入某个节点之后
func (l*list) mvaft(node,after *Node  ){
	tem := l.head
	for tem == after{
		tem = tem.next
	}
	node.next = tem.next
	tem.next= node
}


// 将节点插入某个节点之前
func (l*list) mvbef(node,bef *Node  ){
	tem := l.head
	for tem.next == bef {
		tem = tem.next
	}
	node.next =tem.next
	tem.next = node
}


func (l *list)Range()[]interface{}{
	tem := l.head
	var res []interface{}
	for tem !=nil {
		res= append(res, tem.data)
		tem=tem.next
	}
	return res
}


// 下面是对外接口
// 添加参数
func (l *list)Push (node *Node) {
	if node ==nil {return }
	l.add(node)
}

// 删除节点
func (l *list)Remove(node *Node){
	if  l.head ==nil {return}
	l.rm(node)
}

// 删除倒数第N个节点
func (l *list) RmOrderbyDesc(target int){
	var newList = &list{head: l.head}

	// 设置一个temr
	var temr ,first  =newList.head ,newList.head
	// 双指针： 获取到倒数第target 的位置
	for i:= 0;i<target ;i ++ {
		first = first.next
	}
	// 将快的指针进行穷举，获得慢的指针的位置即删除node的节点
	for first.next !=nil {
		first = first.next
		temr = temr.next
	}
	temr.next =temr.next.next
}



