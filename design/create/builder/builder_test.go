package builder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	director := &Director{builder: &Pig{}}
	director.Construct()
	peppa := director.builder.GetAnimal()
	fmt.Println(peppa)
	assert.Equal(t, peppa.Name, "peppa")
}
