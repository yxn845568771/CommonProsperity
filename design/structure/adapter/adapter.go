package adapter

import "fmt"

// 适配器模式 即让原来不适配的东西，生成他的转化方法另其适配

type worker struct {
}

func (w *worker) inputRoomForMaterial(m material) {
	m.inputRoom()
}

type material interface {
	inputRoom()
}

type cement struct {
}

func (c *cement) inputRoom() {
	fmt.Println("把水泥放进去")
}

type longWood struct {
}

func (l *longWood) obliqueInputRoom() {
	fmt.Println("把长木头斜着放进去")
}

type longWoodAdapter struct {
	longWood *longWood
}

func (l *longWoodAdapter) inputRoom() {
	fmt.Println("先把木头斜过来")
	l.longWood.obliqueInputRoom()
}
