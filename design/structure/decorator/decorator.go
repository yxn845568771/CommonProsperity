package decorator

// 装饰模式 在调用方法时，可以在他的前后加别的代码，如果我们要在原来的代码加上多种新需求，就可以考虑装饰模式

type scale interface {
	getWeight() int64
}

type watermelon struct {
	weight int64
}

type electronicScale struct {
	w watermelon
}

func (e *electronicScale) getWeight() int64 {
	return e.w.weight + 10
}
