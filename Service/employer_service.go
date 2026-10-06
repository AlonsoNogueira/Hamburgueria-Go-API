package service

import (
	"errors"

	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	repository "github.com/alnszzx/HamburgueriaGo/Database/Repository"
	"github.com/google/uuid"
)

type EmployerService struct {
	repository *repository.EmployerRepository
}

func NewEmployerService(repository *repository.EmployerRepository) *EmployerService {
	return &EmployerService{
		repository: repository,
	}
}

func (es *EmployerService) CreateEmployerService(employer *model.Employer) error {
	if employer.Phone == "" || employer.Email == "" || employer.Name == "" {
		return errors.New("Name, phone or Email can´t be nil")
	}

	employer.ID = uuid.New()

	return es.repository.Create(employer)
}

func (es *EmployerService) DeleteEmployerService(id uuid.UUID) error {
	employer, err := es.repository.FindByID(id)
	if err != nil {
		return errors.New("Employer with ID not found")
	}

	return es.repository.Delete(employer.ID)
}
