package house

type visitorInterface interface {
	VisitHouse(houseInterface) string
	VisitBank(bankInterface) string
	VisitFactory(factoryInterface) string
	VisitStreet(streetInterface) string
	Accept(v visitorInterface) string
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
}

type house struct {
	name               string
	FinancialCondition int
	Prudency           bool
}

// Accept visitor at house
func (house house) Accept(v visitorInterface) string {
	return v.VisitHouse(&house)
}

func (h *house) GetName() string {
	return h.name
}

func (h *house) GetFinancialCondition() int {
	return h.FinancialCondition
}
func (h *house) GetPrudency() bool {
	return h.Prudency
}

// Create new house
func NewHouse(name string, FinancialCondition int, Prudency bool) houseInterface {
	return &house{
		name:               name,
		FinancialCondition: FinancialCondition,
		Prudency:           Prudency,
	}
}
