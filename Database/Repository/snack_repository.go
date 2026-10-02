package repository

import (
	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type SnacksRepository struct {
	db *gorm.DB
}

func NewSnacksRepository(db *gorm.DB) *SnacksRepository {
	return &SnacksRepository{
		db: db,
	}
}

func (r *SnacksRepository) GetAll() ([]model.Snacks, error) {
	var snacks []model.Snacks

	err := r.db.Find(&snacks).Error

	return snacks, err
}

func (r *SnacksRepository) Create(snack *model.Snacks) error {
	return r.db.Create(snack).Error
}

func (r *SnacksRepository) FindByID(id uuid.UUID) (*model.Snacks, error) {
	var lanch model.Snacks

	err := r.db.First(&lanch, id).Error
	if err != nil {
		return nil, err
	}

	return &lanch, nil
}

func (r *SnacksRepository) Update(lanch *model.Snacks) error {
	return r.db.Save(lanch).Error
}

func (r *SnacksRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Snacks{}, id).Error
}
