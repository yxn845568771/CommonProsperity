package state

import "testing"

func TestState(t *testing.T) {
	pig := NewPig()
	for {
		pig.eat()
		pig.clean()
		pig.sleep()
	}
}
