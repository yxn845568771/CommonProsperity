package mediator

import "fmt"

// 中介者模式 减少两个或多个模块的依赖，通过中介来平衡，这样好处在于模块不需要考虑其他模块怎么使用，全权又中介解决，减少耦合

type Pig interface {
	Show()
	Sell()
	PermitShow()
}

type BlackPig struct {
	mediator mediator
}

func (p *BlackPig) Show() {
	if !p.mediator.canShow(p) {
		fmt.Println("no show")
		return
	}
	fmt.Println("show")
}

func (p *BlackPig) Sell() {
	fmt.Println("pig sell")
	p.mediator.sellOne()
}

func (p *BlackPig) PermitShow() {
	p.Show()
}

type mediator interface {
	canShow(Pig) bool
	sellOne()
}

type PigMerchant struct {
	length int
	pigs   []Pig
}

func NewPigMerchant(i int) *PigMerchant {
	return &PigMerchant{
		length: i,
		pigs:   nil,
	}
}

func (p *PigMerchant) canShow(pig Pig) bool {
	if len(p.pigs) < p.length {
		p.pigs = append(p.pigs, pig)
		return true
	}
	return false
}

func (p *PigMerchant) sellOne() {
	p.pigs = p.pigs[1:]
}
