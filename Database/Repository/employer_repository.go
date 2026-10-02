package repository

import (
	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type EmployerRepository struct {
	db *gorm.DB
}

func NewEmployerRepository(db *gorm.DB) *EmployerRepository {
	return &EmployerRepository{
		db: db,
	}
}

func (r *EmployerRepository) Create(employer *model.Employer) error {
	return r.db.Create(employer).Error
}

func (r *EmployerRepository) FindAll() ([]model.Employer, error) {
	var employers []model.Employer

	err := r.db.Find(&employers).Error

	return employers, err
}

func (r *EmployerRepository) FindByID(id uuid.UUID) (*model.Employer, error) {
	var employer *model.Employer

	err := r.db.First(&employer, id).Error
	if err != nil {
		return nil, err
	}

	return employer, nil
}

func (r *EmployerRepository) Update(employer *model.Employer) error {
	return r.db.Save(employer).Error
}

func (r *EmployerRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Employer{}, id).Error
}
