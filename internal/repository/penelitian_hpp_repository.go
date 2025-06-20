package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianHPPRepository struct {
	Repository[entity.PenelitianHPP]
	Log *logrus.Logger
}

func NewPenelitianHPPRepository(log *logrus.Logger) *PenelitianHPPRepository {
	return &PenelitianHPPRepository{Log: log}
}

func (r *PenelitianHPPRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianHPP, error) {
	var result []entity.PenelitianHPP
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
} 