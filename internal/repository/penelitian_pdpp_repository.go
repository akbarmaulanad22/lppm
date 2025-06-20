package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianPDPPRepository struct {
	Repository[entity.PenelitianPDPP]
	Log *logrus.Logger
}

func NewPenelitianPDPPRepository(log *logrus.Logger) *PenelitianPDPPRepository {
	return &PenelitianPDPPRepository{Log: log}
}

func (r *PenelitianPDPPRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianPDPP, error) {
	var result []entity.PenelitianPDPP
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
