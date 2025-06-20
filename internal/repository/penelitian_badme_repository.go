package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianBADMERepository struct {
	Repository[entity.PenelitianBADME]
	Log *logrus.Logger
}

func NewPenelitianBADMERepository(log *logrus.Logger) *PenelitianBADMERepository {
	return &PenelitianBADMERepository{Log: log}
}

func (r *PenelitianBADMERepository) FindAll(tx *gorm.DB) ([]entity.PenelitianBADME, error) {
	var result []entity.PenelitianBADME
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
