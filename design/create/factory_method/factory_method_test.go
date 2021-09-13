package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenMilkFactory(t *testing.T) {
	factory1, err := GenMilkFactory(MilkFactoryCow)
	assert.NoError(t, err)
	assert.Equal(t, factory1.SellMilk().Taste, "味道还行")
	factory2, err := GenMilkFactory(MilkFactoryPig)
	assert.NoError(t, err)
	assert.Equal(t, factory2.SellMilk().Taste, "狗都不喝")
	_, err = GenMilkFactory(MilkFactoryUnKnow)
	assert.Error(t, err)
}
