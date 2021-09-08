package singleton

import "sync"

// 懒汉式 好处在于没有使用new方法的话不占用资源

// 对象私有化
type pig struct {
	name   string
	weight float64
	price  float64
}

func (p *pig) SetPrice(price float64) {
	p.price = price
}

func (p *pig) GetPrice() float64 {
	return p.price
}

// 使用once保证在多线程的情况下只有实例
var (
	once  sync.Once
	peppa *pig
)

// 对外仅暴露一个生成对象的接口，控制其生成方式
func NewPig() *pig {
	once.Do(func() {
		peppa = &pig{
			name:   "佩琪",
			weight: 1800,
			price:  20,
		}
	})
	return peppa
}
