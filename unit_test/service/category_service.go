package service

import (
	"errors"
	"unit_test/entity"
	"unit_test/repository"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
}

func (service CategoryServiceImpl) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	}

	return category, nil
}