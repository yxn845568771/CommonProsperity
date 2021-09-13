package observer

import "fmt"

// 观察者模式	该模式适用于需要能够及时知道别处变动的场景，比如说公告改了，客户要立马知道

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type Pig struct {
	Weight int64
	Price  int64
}

type pigGuy struct {
	observers []observer
	pig       *Pig
}

func (p *pigGuy) register(observer observer) {
	p.observers = append(p.observers, observer)
}

func (p *pigGuy) deregister(observer observer) {
	for i, o := range p.observers {
		if o.getName() == observer.getName() {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
		}
	}
}

func (p *pigGuy) updatePig(pig *Pig) {
	p.pig = pig
	p.notifyAll()
}

func (p *pigGuy) notifyAll() {
	for _, o := range p.observers {
		o.receive(p.pig)
	}
}

type observer interface {
	receive(pig *Pig)
	getName() string
}

type customer struct {
	name string
}

func (c *customer) getName() string {
	return c.name
}

func (c *customer) receive(pig *Pig) {
	fmt.Println(*pig)
}
