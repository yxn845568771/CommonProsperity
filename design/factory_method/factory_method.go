package factory_method

import "errors"

// 工厂方法模式
// 将两个或多个具有相同功能的方法抽成interface 他们分别实现这个interface 而其他对象想用他们时只需调用一个new方法即可(加上筛选条件)
// 好处：如果多了一个功能相似的对象，只需要实现方法即可,在改动原来的new接口即可，减少代码耦合
// 坏处：每次增加都要改变new方法及使用new方法的地方

// 卖奶的工厂
type MilkFactory interface {
	SellMilk() *Milk
}

// 奶
type Milk struct {
	Price float64
	Taste string
}

// 牛工厂
type CowFactory struct {
	mick *Milk
}

func (c *CowFactory) SellMilk() *Milk {
	return c.mick
}

// 猪工厂
type PigFactory struct {
	mick *Milk
}

func (c *PigFactory) SellMilk() *Milk {
	return c.mick
}

// 动物类型
type AnimalType uint

const (
	MilkFactoryUnKnow AnimalType = iota
	MilkFactoryCow
	MilkFactoryPig
)

// 得到一个买奶的工厂
func GenMilkFactory(animal AnimalType) (MilkFactory, error) {
	switch animal {
	case MilkFactoryCow:
		return &CowFactory{mick: &Milk{
			Price: 100,
			Taste: "味道还行",
		}}, nil
	case MilkFactoryPig:
		return &PigFactory{mick: &Milk{
			Price: 0.001,
			Taste: "狗都不喝",
		}}, nil
	default:
		return nil, errors.New("这边建议你和豆浆")
	}
}
