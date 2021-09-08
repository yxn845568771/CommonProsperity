package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPig(t *testing.T) {
	peppa1 := NewPig()
	peppa2 := NewPig()
	peppa1.SetPrice(peppa1.GetPrice() - 1)
	assert.Equal(t,peppa1.GetPrice(), peppa2.GetPrice())
}
