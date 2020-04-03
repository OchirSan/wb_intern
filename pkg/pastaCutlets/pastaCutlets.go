package pastaCutlets

type dishInterface interface {
	SetMeat(string)
	SetGarnish(string)
	ShowDish()
}

type PastaCutlets interface {
	FryMeat()
	MakeSideDish()
	ShowDish()
}

type pastaCutlets struct {
	dish dishInterface
}

func (p *pastaCutlets) FryMeat() {
	p.dish.SetMeat("beef patties")
}

func (p *pastaCutlets) ShowDish() {
	p.dish.ShowDish()
}

func (p *pastaCutlets) MakeSideDish() {
	p.dish.SetGarnish("pasta")
}

// Create new pasta with Cutlets
func NewPastaCutlets(dish dishInterface) PastaCutlets {
	return &pastaCutlets{
		dish,
	}
}
