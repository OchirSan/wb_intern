package builder

type dishBuilder interface {
	FryMeat()
	MakeSideDish()
}

// Director interface to work with builder
type Director interface {
	BuildDish()
}

type director struct {
	builder dishBuilder
}

// Concrete builder build new dish
func (d *director) BuildDish() {
	d.builder.FryMeat()
	d.builder.MakeSideDish()
}

// Create new director, whose leads the builder
func NewDirector(builder dishBuilder) Director {
	return &director{
		builder: builder,
	}
}
