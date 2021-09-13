package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterator(t *testing.T) {
	pig := &Pig{Weight: 200}
	cow := &Cow{Weight: 1000}
	iterator := &Iterator{
		index:     0,
		Container: Container{},
	}
	iterator.Add(pig)
	iterator.Add(cow)
	assert.Equal(t,GetSumWeight(iterator),float64(1200))
}

func GetSumWeight(iterator *Iterator) (sum float64) {
	for {
		if !iterator.HasNext() {
			break
		}
		sum += iterator.Next().GetWeight()
	}
	return
}
