package visitor

// 访问者模式 ,原来的爱心代码怕被改出事，所以就改用这个模式，只需要在每个对象上实现accept，然后实现visitor就可以使用私有对象，
// 好处:只需要实现accept就可以修改成任意功能
// 坏处:接口名不直观

type animal interface {
	getName() string
	accept(visitor)
}

type pig struct {
	name string
}

type visitor interface {
	visitForPig(p *pig)
	visitForCow(c *cow)
}

func (p *pig) getName() string {
	return p.name
}

func (p *pig) accept(v visitor) {
	v.visitForPig(p)
}

type cow struct {
	name string
}

func (c *cow) getName() string {
	return c.name
}

func (c *cow) accept(v visitor) {
	v.visitForCow(c)
}

type nameUpdater struct {
	pigName string
	cowName string
}

func (n *nameUpdater) visitForPig(p *pig) {
	p.name = n.pigName
}

func (n *nameUpdater) visitForCow(c *cow) {
	c.name = n.cowName
}
