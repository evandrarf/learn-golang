package repository

import "unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category 
}