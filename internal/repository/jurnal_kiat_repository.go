package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalKIATRepository struct {
	Repository[entity.JurnalKIAT]
	Log *logrus.Logger
}

func NewJurnalKIATRepository(log *logrus.Logger) *JurnalKIATRepository {
	return &JurnalKIATRepository{Log: log}
}

func (r *JurnalKIATRepository) FindAll(tx *gorm.DB) ([]entity.JurnalKIAT, error) {
	var result []entity.JurnalKIAT
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
