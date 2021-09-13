package memento

// 备忘录模式 感觉没啥东西，就是记录改动前的样子，如果加上装饰模式的话就可以自动加快照，岂不美哉

type pig struct {
	weight int64
	price  int64
}

func (p *pig) sell(meat int64) {
	p.weight -= meat
	p.price += meat / 10
}

type Memento struct {
	Pig pig
}

func NewMemento(p *pig) Memento {
	return Memento{Pig: *p}
}

type recorder struct {
	Mementos []Memento
}

func (r *recorder) add(memento Memento) {
	r.Mementos = append(r.Mementos, memento)
}


