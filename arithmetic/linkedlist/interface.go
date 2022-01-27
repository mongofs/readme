package linkedlist

type LinkedList interface {
	MoveToFront(node *Node)        // 移动node到链表之前
	MoveToBack(node *Node)         // 移动node到链表尾部
	MoveBefore(node, target *Node) // 移动node 到 target之前
	MoveAfter(node, target *Node)  // 移动node到 target 之后

	PushFront(node *Node)               //添加节点到链表头部
	PushBack(node *Node)                //添加节点到链表尾部
	PushBefore(node, target *Node)       // 添加到target之前
	PushAfter(node, target *Node)        // 添加到target之后

	RemoveNode(node *Node)               //移除Node
	RemoveLastKthNode(lastKth int) *Node // 移除倒数第k个Node
	RemoveLast() *Node                   // 移除链表尾部Node
	RemoveHead() *Node                   // 移除链表头部Node

	Reverse() // 反转整个链表
	Sort()    //排序整个链表
}
