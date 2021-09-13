package stratrgy

import "fmt"

// 策略模式 我们常常写代码的时候会发现里面有着几乎一样的逻辑，这时候就可以考虑将其部分抽成方法，其他来实现他，到时候只需要将实现放到结构体里即可

type reward interface {
	issue(singleT int64) int64
}

type order struct {
	power int64
	ratio int64
}

type account struct {
	balance int64
}

type immediately struct {
	order order
	acc   *account
}

func (i *immediately) issue(singleT int64) int64 {
	amount := i.order.power * i.order.ratio * singleT / 4
	i.acc.balance += amount
	return amount
}

type linear struct {
	order order
	lin   *account
}

func (i *linear) issue(singleT int64) int64 {
	amount := i.order.power * i.order.ratio * singleT * 3 / 4
	i.lin.balance += amount
	return amount
}

type rewardHelper struct {
	reward reward
}

func (r *rewardHelper) issue(singleT int64) {
	amount := r.reward.issue(singleT)
	fmt.Println(amount)
}
