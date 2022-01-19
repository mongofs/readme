// 观察者模式，是由资源本体结构提供接口注册，

package observer

type Project struct {
	obs []Observer
	ctx string
}

func (p *Project) Attach(o Observer){
	p.obs = append(p.obs, o)
}

// 被监控的操作
func (p *Project)UpdateContext (ctx string){
	p.ctx =ctx
	p.notify()
}

func (p *Project)notify (){
	for _,o := range  p.obs {
		o.Update(p)
	}
}


// 注册结果
type Observer  interface {
	Update(project *Project)
}