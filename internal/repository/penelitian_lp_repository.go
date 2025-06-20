package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianLPRepository struct {
	Repository[entity.PenelitianLP]
	Log *logrus.Logger
}

func NewPenelitianLPRepository(log *logrus.Logger) *PenelitianLPRepository {
	return &PenelitianLPRepository{Log: log}
}

func (r *PenelitianLPRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianLP, error) {
	var result []entity.PenelitianLP
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
