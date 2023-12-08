package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	barRepository *BarRepository
}

func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{barRepository: barRepository}
}