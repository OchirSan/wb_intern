package factory

type visitorInterface interface {
	VisitHouse(houseInterface) string
	VisitBank(bankInterface) string
	VisitFactory(factoryInterface) string
	VisitStreet(streetInterface) string
}

type visitor struct{}

type houseInterface interface {
	GetName() string
	GetFinancialCondition() int
	GetPrudency() bool
}

type bankInterface interface {
	GetName() string
	GetIsHadIndurance() bool
}

type factoryInterface interface {
	GetName() string
	GetIsHadDanger() bool
}

type streetInterface interface {
	GetName() string
	GetHouse() houseInterface
	GetBank() bankInterface
	GetFactory() factoryInterface
	Accept(v visitorInterface) string
}

type factory struct {
	name        string
	IsHadDanger bool
}

// Accept visitor at factory
func (factory factory) Accept(v visitorInterface) string {
	return v.VisitFactory(&factory)
}

func (f *factory) GetName() string {
	return f.name
}

func (f *factory) GetIsHadDanger() bool {
	return f.IsHadDanger
}

// Create new factory
func NewFactory(name string, IsHadDanger bool) factoryInterface {
	return &factory{
		name:        name,
		IsHadDanger: IsHadDanger,
	}
}
