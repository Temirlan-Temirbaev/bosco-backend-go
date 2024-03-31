package service

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
)

type CategoryService struct {
	repository repository.Category
}

func NewCategoryService(repository repository.Category) *CategoryService {
	return &CategoryService{repository: repository}
}

func (service *CategoryService) Create(category model.Category) (int, error) {
	return service.repository.Create(category)
}

func (service *CategoryService) Update(id int, category model.Category) error {
	return service.repository.Update(id, category)
}

func (service *CategoryService) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *CategoryService) GetAll() ([]model.Category, error) {
	return service.repository.GetAll()
}

func (service *CategoryService) GetById(id int) (model.Category, error) {
	return service.repository.GetById(id)
}
