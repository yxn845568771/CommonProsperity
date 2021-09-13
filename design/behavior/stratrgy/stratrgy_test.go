package stratrgy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	order := order{
		power: 100,
		ratio: 1,
	}
	lin := &account{balance: 0}
	helper := rewardHelper{reward: &linear{
		order: order,
		lin:   lin,
	}}
	helper.issue(1)
	fmt.Println(lin)
}
