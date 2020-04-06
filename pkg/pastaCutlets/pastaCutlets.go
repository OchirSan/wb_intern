package pasta_cutlets

type dishInterface interface {
	SetMeat(string)
	SetGarnish(string)
	ShowDish()
	GetMeat() int
	GetGarnish() int
	IsItDish() string
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
