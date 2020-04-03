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

func (d *director) BuildDish() {
	d.builder.FryMeat()
	d.builder.MakeSideDish()
}

// Create new director
func NewDirector(builder dishBuilder) Director {
	return &director{
		builder: builder,
	}
}
