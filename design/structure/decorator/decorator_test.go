package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	scale := &electronicScale{w: watermelon{weight: 10}}
	weight := scale.getWeight()
	if weight > 10 {
		fmt.Println("你这瓜保熟吗")
		fmt.Println("你tm劈我瓜是吧")
	}
}
