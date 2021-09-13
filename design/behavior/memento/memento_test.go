package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	pig := &pig{
		weight: 100,
		price:  10,
	}
	
	recorder := recorder{}
	recorder.add(NewMemento(pig))
	
	pig.sell(10)
	recorder.add(NewMemento(pig))
	
	pig.sell(20)
	recorder.add(NewMemento(pig))
	
	for _, memento := range recorder.Mementos {
		fmt.Println(memento)
	}
}
