package bridge

import "fmt"

// 桥接模式，上下层关系可独立开发 不需要理其它层的开发

type animal interface {
	taste()
}

type pig struct {
}

func (p *pig) taste() {
	fmt.Println("狗都不吃")
}

type cow struct {
}

func (c *cow) taste() {
	fmt.Println("狗喜欢吃")
}

type plant interface {
	evaluation()
	setAnimal(animal)
}

type canPlant struct {
	animal animal
}

func (p *canPlant) evaluation() {
	fmt.Println("罐头古法制造")
	p.animal.taste()
}

func (p *canPlant) setAnimal(animal animal) {
	p.animal = animal
}

type sausagePlant struct {
	animal animal
}

func (p *sausagePlant) evaluation() {
	fmt.Println("腊肠古法制造")
	p.animal.taste()
}

func (p *sausagePlant) setAnimal(animal animal) {
	p.animal = animal
}
