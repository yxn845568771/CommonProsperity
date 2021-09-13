package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPig_Clone(t *testing.T) {
	peppa := &Pig{
		Name:   "佩琪",
		Price:  123,
		Weight: 123,
		Sons: []*Pig{
			{
				Name:   "佩琪大儿子",
				Price:  123,
				Weight: 123,
			},
			{
				Name:   "佩琪二女儿",
				Price:  123,
				Weight: 123,
			},
		},
	}
	copyPeppa := peppa.Clone()
	copyPeppa.Sons[0].Price = 1
	assert.NotEmpty(t, copyPeppa.Sons[0].Price, peppa.Sons[0].Price)
}