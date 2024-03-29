package service

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
)

type ContactService struct {
	repository repository.Contact
}

func NewContactService(repository repository.Contact) *ContactService {
	return &ContactService{repository: repository}
}

func (service *ContactService) Create(contact model.Contact) (int, error) {
	return service.repository.Create(contact)
}

func (service *ContactService) GetAll() ([]model.Contact, error) {
	return service.repository.GetAll()
}
