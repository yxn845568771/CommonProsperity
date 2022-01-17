package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	scale := &electronicScale{w: newProxy()}
	fmt.Println(scale.getWeight())
}
