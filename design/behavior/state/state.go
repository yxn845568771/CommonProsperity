package state

import (
	"fmt"
	"time"
)

// 状态模式 将每一个状态都写成一个类独立的处理，这样可以很容易的控制状态的改变，但是缺点也明细，实现了太多无用方法

type state interface {
	eat()
	sleep()
	clean()
}

type pig struct {
	isHungry    state
	isHappiness state
	isDirty     state
	
	currentState state
}

func (p pig) eat() {
	p.isHungry.eat()
}

func (p pig) sleep() {
	p.isHappiness.sleep()
}

func (p pig) clean() {
	p.isDirty.clean()
}

func NewPig() *pig {
	pig := &pig{}
	pig.isDirty = &isDirty{pig: pig}
	pig.isHungry = &isHungry{pig: pig}
	pig.isHappiness = &isHappiness{pig: pig}
	fmt.Println("幸福猪猪刚刚出生")
	pig.currentState = pig.isDirty
	return pig
}

type isHungry struct {
	pig *pig
}

func (i isHungry) eat() {
	fmt.Println("猪吃得很开心")
	fmt.Println("但是它吃的很脏，需要洗澡")
	i.pig.currentState = i.pig.isDirty
}

func (i isHungry) sleep() {
	panic("implement me")
}

func (i isHungry) clean() {
	panic("implement me")
}

type isHappiness struct {
	pig *pig
}

func (i isHappiness) eat() {
	panic("implement me")
}

func (i isHappiness) sleep() {
	fmt.Println("猪现在很开心，他要睡觉了")
	time.Sleep(10 * time.Second)
	fmt.Println("猪现在醒了，他饿了")
	i.pig.currentState = i.pig.isHungry
}

func (i isHappiness) clean() {
	panic("implement me")
}

type isDirty struct {
	pig *pig
}

func (i isDirty) eat() {
	panic("implement me")
}

func (i isDirty) sleep() {
	panic("implement me")
}

func (i isDirty) clean() {
	fmt.Println("猪吃完东西很脏，需要洗洗澡")
	fmt.Println("猪现在洗的很干净，他们幸福")
	i.pig.currentState = i.pig.isHappiness
}
