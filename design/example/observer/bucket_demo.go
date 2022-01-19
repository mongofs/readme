// 在这里解释说明一下为什么bucket不使用观察者模式，因为本身golang 的interface 类型存在
// 更多的空间存放值和类型的指针，在维护用户的连接状态的对象中需要更小的内存开销，所以不适合
// 这里的使用，第二个原因是由于我们在bucket存放用户的连接信息，在操作过程中存在用户单条信息
// 查找，这里也不适合数组存放，相反使用互斥锁更加适合。但是互斥锁如果存放cli的Clienter类型
// 难免有点脱裤子放屁的嫌疑

package observer

import "im/client"

type Bucket struct {
	clis []client.Clienter
}


func (b *Bucket)Add(cli client.Clienter){
	b.clis= append(b.clis,cli)
}


func (b *Bucket)BroadCast(message []byte){
	for _,v := range b.clis{
		v.Send(message)
	}
}

