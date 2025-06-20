package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalTAJBRepository struct {
	Repository[entity.JurnalTAJB]
	Log *logrus.Logger
}

func NewJurnalTAJBRepository(log *logrus.Logger) *JurnalTAJBRepository {
	return &JurnalTAJBRepository{Log: log}
}

func (r *JurnalTAJBRepository) FindAll(tx *gorm.DB) ([]entity.JurnalTAJB, error) {
	var result []entity.JurnalTAJB
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
} 