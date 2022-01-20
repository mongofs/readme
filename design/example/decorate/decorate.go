// 装饰器本身是 针对类的内容进行扩展
package main

type User interface {
	GetUserInfo()string
}

type store struct {}


func (s *store) GetUserInfo()string{
	return "steven"
}


// 接下来是装饰器的结构，这里的需求是把之前的返回的名字变成name.age这样的形式返回
// 其实装饰器模式其实还是比较简单，就是将之前的方法在不变动的情况下进行重构，但也可以考虑
// 装饰成本的高低，如果说层级过于深入，我认为是没有太大的必要的，不过在基础封装包内可以进行
// 使用，但是在业务层面上我还是不太建议使用装饰器模式进行包装。中间件一般来说也是装饰器的宠儿
// 装饰器模式是完全遵守开闭原则
type UserDecorate struct {
	User
}

// 这里存在一个方法对原来的方法进行装饰
// 比如我要在string后面加上age ： name.age
func (u *UserDecorate) GetUserInfo()string{
	return u.User.GetUserInfo()+"."+"28"
}
