package bank

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

// Accept visitor at bank
func (bank bank) Accept(v visitorInterface) string {
	return v.VisitBank(&bank)
}

type bank struct {
	name           string
	IsHadInsurance bool
}

func (b *bank) GetName() string {
	return b.name
}

func (b *bank) GetIsHadIndurance() bool {
	return b.IsHadInsurance
}

// Create new bank
func NewBank(name string, IsHadInsurance bool) bankInterface {
	return &bank{
		name:           name,
		IsHadInsurance: IsHadInsurance,
	}
}
