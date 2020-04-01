package factory

import (
	"fmt"
)

type factoryInterface interface {
	VisitFactory() string
}

type factory struct {
	name        string
	IsHadDanger bool
}

// Accept visitor in factory
func Accept(f factoryInterface) string {
	return f.VisitFactory()
}

// Visitor visit factory
func (factory factory) VisitFactory() string {
	res := fmt.Sprintln("visiting", factory.name, "house and offer fire and flood insurance")
	if factory.IsHadDanger {
		res += fmt.Sprintln("accept visitor")
	} else {
		res += fmt.Sprintln("kick off visitor")
	}
	return res
}

// Create new factory
func NewFactory(name string, IsHadDanger bool) factoryInterface {
	return &factory{
		name:        name,
		IsHadDanger: IsHadDanger,
	}
}
