package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianSTPRepository struct {
	Repository[entity.PenelitianSTP]
	Log *logrus.Logger
}

func NewPenelitianSTPRepository(log *logrus.Logger) *PenelitianSTPRepository {
	return &PenelitianSTPRepository{Log: log}
}

func (r *PenelitianSTPRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianSTP, error) {
	var result []entity.PenelitianSTP
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
