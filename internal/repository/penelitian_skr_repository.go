package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianSKRRepository struct {
	Repository[entity.PenelitianSKR]
	Log *logrus.Logger
}

func NewPenelitianSKRRepository(log *logrus.Logger) *PenelitianSKRRepository {
	return &PenelitianSKRRepository{Log: log}
}

func (r *PenelitianSKRRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianSKR, error) {
	var result []entity.PenelitianSKR
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
