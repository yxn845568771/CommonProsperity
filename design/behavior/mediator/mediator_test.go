package mediator

import "testing"

func TestMediator(t *testing.T) {
	chant := NewPigMerchant(2)
	blackPig := &BlackPig{mediator: chant}
	blackPig2 := &BlackPig{mediator: chant}
	blackPig3 := &BlackPig{mediator: chant}
	blackPig.Show()
	blackPig2.Show()
	blackPig3.Show()
	chant.sellOne()
	blackPig3.Show()
}
