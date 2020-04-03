package pilaf

type dishInterface interface {
	SetMeat(string)
	SetGarnish(string)
	ShowDish()
}

type Pilaf interface {
	FryMeat()
	MakeSideDish()
	ShowDish()
}

type pilaf struct {
	dish dishInterface
}

func (p *pilaf) FryMeat() {
	p.dish.SetMeat("chicken meat")
}

func (p *pilaf) ShowDish() {
	p.dish.ShowDish()
}

func (p *pilaf) MakeSideDish() {
	p.dish.SetGarnish("rice")
}

// Create new pilaf
func NewPilaf(dish dishInterface) Pilaf {
	return &pilaf{
		dish,
	}
}
