package observer

import "testing"

func TestObserver(t *testing.T) {
	c1 := &customer{name: "123"}
	c2 := &customer{name: "456"}
	c3 := &customer{name: "789"}
	
	guy := pigGuy{}
	guy.register(c1)
	guy.register(c2)
	guy.register(c3)
	guy.updatePig(&Pig{
		Weight: 100,
		Price:  10,
	})
	guy.deregister(c3)
	guy.updatePig(&Pig{
		Weight: 200,
		Price:  20,
	})
}
