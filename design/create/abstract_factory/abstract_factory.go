package abstract_factory

// 抽象工厂模式：在工厂方法的基础上把工厂也抽象了，避免了工厂在new时不必要的参数输入
// 好处：利于需求的扩建，只需要再实现工厂和产品即可增加功能，解耦进一步加强
// 坏处：工厂也被抽成interface,代码工作量加重,如果工厂种类不多则不如工厂方法

// 当铺
type Pawn interface {
	CreateGoods() Goods
}

// 商品
type Goods interface {
	GetPrice() float64
}

// 肉
type Meat struct {
	Price float64
}

func (m *Meat) GetPrice() float64 {
	return m.Price
}

// 猪肉档
type PigPawn struct {
	meat *Meat
}

func (p *PigPawn) CreateGoods() Goods {
	p.meat = &Meat{Price: 100}
	return p.meat
}

// 水果
type Fruit struct {
	Price float64
}

func (f *Fruit) GetPrice() float64 {
	return f.Price
}

// 水果档
type FruitPawn struct {
	fruit *Fruit
}

func (f *FruitPawn) CreateGoods() Goods {
	f.fruit = &Fruit{Price: 10}
	return f.fruit
}
