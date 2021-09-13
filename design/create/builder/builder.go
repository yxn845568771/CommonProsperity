package builder

// 建造者模式：面对参数多的结构体初始化，一个个赋值容易搞混，使用建造者模式来解决

type Animal struct {
	Name   string
	Price  float64
	Weight float64
}

type Builder interface {
	SetName() Builder
	SetPrice() Builder
	SetWeight() Builder
	GetAnimal() Animal
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.SetName().SetPrice().SetWeight()
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

type Pig struct {
	animal Animal
}

func (p *Pig) SetName() Builder {
	p.animal.Name = "peppa"
	return p
}

func (p *Pig) SetPrice() Builder {
	p.animal.Price = 100
	return p
}

func (p *Pig) SetWeight() Builder {
	p.animal.Weight = 10000
	return p
}

func (p *Pig) GetAnimal() Animal {
	return p.animal
}
