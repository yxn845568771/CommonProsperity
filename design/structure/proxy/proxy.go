package proxy

type scale interface {
	getWeight() int64
}

type watermelon struct {
	weight int64
}

type electronicScale struct {
	w watermelon
}

func newProxy() watermelon {
	return watermelon{weight: 10}
}

func (e *electronicScale) getWeight() int64 {
	return e.w.weight + 10
}
