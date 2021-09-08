package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMeat_GetPrice(t *testing.T) {
	pigPawn := &PigPawn{}
	goods := pigPawn.CreateGoods()
	assert.Equal(t,goods.GetPrice(),float64(100))
}
