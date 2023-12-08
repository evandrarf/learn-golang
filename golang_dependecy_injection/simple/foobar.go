package simple

type FooBarService struct {
	fooService *FooService
	barService *BarService
}

func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{fooService: fooService, barService: barService}
}