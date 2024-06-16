package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(foo *FooService, bar *BarService) *FooBarService {
	return &FooBarService{
		FooService: foo,
		BarService: bar,
	}
}