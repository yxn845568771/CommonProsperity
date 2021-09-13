package iterator

// 迭代器模式 : 当我们需要便利一堆数据，我们要考虑让他们抽象成同一个结构，方便便利

// 实现一个迭代器，里面存放下标及数组
type Iterator struct {
	index int
	Container
}

func (i *Iterator) Next() Animal {
	visitor := i.list[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}

// 抽象数组
type Container struct {
	list []Animal
}

func (c *Container) Add(animal Animal) {
	c.list = append(c.list, animal)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

type Animal interface {
	GetWeight() float64
}

type Pig struct {
	Weight float64
}

func (p *Pig) GetWeight() float64 {
	return p.Weight
}

type Cow struct {
	Weight float64
}

func (c *Cow) GetWeight() float64 {
	return c.Weight
}
