// 适配器模式是因为老的接口不满足新的接口的被需求的状态，举个简单例子，我已经实现了一个older 的类
// 但是启用一个A服务需要注册一个Younger的接口，但是实际上我们并不需要再次实现一个新的younger 的实现，
// 我们可以将新的实现内部包裹旧的接口，就可以完成一个适配器的基本需求
package main

import "fmt"

type  Older interface {
	SetUserName (name string)error
}

type older  struct {
	name string
}

func (o * older) SetUserName(name string) error {
	o.name =name
	fmt.Println(name)
	return nil
}
// 上面是老的接口 和老的具体实现

// 我们有一个新的需求的接口被需要
type Younger interface {
	SetName (name string)error
}



// 使用适配器模式去进行新老接口的对接
type younger struct {
	Older
}

func (y *younger)SetName(name string) error{
	return y.Older.SetUserName(name)
}



func main (){
	// 这段代码的意义在于：
	// 我有一个结构体需要用户注册一个Younger接口
	// 我们在实际情况中我们已经存在一个demo的使用方式了
	// 那么我们就可以使用一个新的结构体包裹demo，然后进行注册使用

	// 实际应用场景可以体会im项目中的 options.Validate选项
	a := func(y Younger) {
		y.SetName("我是个帅哥")
	}
	youth := &younger{&older{}}
	a(youth)
}
