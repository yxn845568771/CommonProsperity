package prototype

type Animal struct {
	Name   string
	Price  float64
	Weight float64
}

type Pig struct {
	Name   string
	Price  float64
	Weight float64
	Sons   []*Pig
}

func (p *Pig) Clone() *Pig {
	copyPig := &Pig{
		Name:   p.Name,
		Price:  p.Price,
		Weight: p.Weight,
	}
	pigs := make([]*Pig, len(p.Sons))
	for i, son := range p.Sons {
		pigs[i] = son.Clone()
	}
	copyPig.Sons = pigs
	return copyPig
}
