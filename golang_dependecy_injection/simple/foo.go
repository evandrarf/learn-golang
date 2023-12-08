package simple

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	fooRepository *FooRepository
}

func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{fooRepository: fooRepository}
}