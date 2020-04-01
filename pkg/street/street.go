package street

type Visitor interface {
	VisitHouse(HouseInterface) string
	VisitBank(BankInterface) string
	VisitFactory(FactoryInterface) string
	VisitStreet(StreetInterface) string
}

type visitor struct{}

type HouseInterface interface {
	GetName() string
	GetFinancialCondition() int
	GetPrudency() bool
}

type BankInterface interface {
	GetName() string
	GetIsHadIndurance() bool
}

type FactoryInterface interface {
	GetName() string
	GetIsHadDanger() bool
}

type StreetInterface interface {
	GetName() string
	GetHouse() HouseInterface
	GetBank() BankInterface
	GetFactory() FactoryInterface
	Accept(v Visitor) string
}

type street struct {
	name    string
	house   HouseInterface
	bank    BankInterface
	factory FactoryInterface
}

// Accept visitor at street
func (s street) Accept(v Visitor) string {
	return v.VisitStreet(&s)
}

func (s *street) GetName() string {
	return s.name
}

func (s *street) GetHouse() HouseInterface {
	return s.house
}

func (s *street) GetBank() BankInterface {
	return s.bank
}

func (s *street) GetFactory() FactoryInterface {
	return s.factory
}

// Create new street
func NewStreet(name string, house HouseInterface, factory FactoryInterface, bank BankInterface) StreetInterface {
	return &street{
		name:    name,
		house:   house,
		factory: factory,
		bank:    bank,
	}
}
