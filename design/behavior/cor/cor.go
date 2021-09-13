package cor

// 责任链模式，将长而复杂的逻辑拆分成一步一步具体的动作，链可以依照我的想法创建顺序，代码也能实现解耦，只需要改变set时的对象

type Pig struct {
	Name string
}

func (p *Pig) clear() {}

type Department interface {
	Execute(*Pig)
	SetNext(Department)
}

type Farms struct {
	next Department
}

func (f *Farms) Execute(pig *Pig) {
	pig.clear()
	// 把猪拉出来买
	if f.next != nil {
		f.next.Execute(pig)
	}
}

func (f *Farms) SetNext(d Department) {
	f.next = d
}

type Slaughterhouse struct {
	next Department
}

func (f *Slaughterhouse) Execute(pig *Pig) {
	// 把猪分尸
	if f.next != nil {
		f.next.Execute(pig)
	}
}

func (f *Slaughterhouse) SetNext(d Department) {
	f.next = d
}

type Market struct {
	next Department
}

func (f *Market) Execute(pig *Pig) {
	// 把猪卖给客户
	if f.next != nil {
		f.next.Execute(pig)
	}
}

func (f *Market) SetNext(d Department) {
	f.next = d
}

type Person struct {
	next Department
}

func (f *Person) Execute(pig *Pig) {
	// 把猪吃了
	if f.next != nil {
		f.next.Execute(pig)
	}
}

func (f *Person) SetNext(d Department) {
	f.next = d
}
