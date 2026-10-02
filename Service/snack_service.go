package service

import (
	"errors"

	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	repository "github.com/alnszzx/HamburgueriaGo/Database/Repository"
	uuid "github.com/google/uuid"
)

type SnacksService struct {
	repository *repository.SnacksRepository
}

func NewSnackService(repository *repository.SnacksRepository) *SnacksService {
	return &SnacksService{
		repository: repository,
	}
}

func (sns *SnacksService) CreateSnackService(snack *model.Snacks) error {
	if snack.Price < 0 {
		return errors.New("Price can not null")
	}

	snack.ID = uuid.New()

	return sns.repository.Create(snack)
}

func (sns *SnacksService) GetSnackByIdService(id uuid.UUID) (*model.Snacks, error) {
	snack, err := sns.repository.FindByID(id)

	if err != nil {
		return nil, errors.New("Snack not found")
	}

	return snack, nil
}

func (sns *SnacksService) UpdateSnackService(id uuid.UUID, snack *model.Snacks) error {
	existingSnack, err := sns.repository.FindByID(id)
	if err != nil {
		return errors.New("Snack not found")
	}

	if snack.Price < 0 {
		return errors.New("Price can not negative")
	}

	existingSnack.Name = snack.Name
	existingSnack.Description = snack.Description
	existingSnack.Price = snack.Price

	return sns.repository.Update(existingSnack)
}

func (sns *SnacksService) DeleteSnackService(id uuid.UUID) error {
	snack, err := sns.repository.FindByID(id)
	if err != nil {
		return errors.New("Snack not found")
	}

	return sns.repository.Delete(snack.ID)
}
