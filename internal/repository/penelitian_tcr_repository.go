package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianTCRRepository struct {
	Repository[entity.PenelitianTCR]
	Log *logrus.Logger
}

func NewPenelitianTCRRepository(log *logrus.Logger) *PenelitianTCRRepository {

	return &PenelitianTCRRepository{Log: log}

}

func (r *PenelitianTCRRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianTCR, error) {
	var PenelitianTCR []entity.PenelitianTCR
	if err := tx.Find(&PenelitianTCR).Error; err != nil {
		return nil, err
	}
	return PenelitianTCR, nil
}
