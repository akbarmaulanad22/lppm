package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalTeknoisRepository struct {
	Repository[entity.JurnalTeknois]
	Log *logrus.Logger
}

func NewJurnalTeknoisRepository(log *logrus.Logger) *JurnalTeknoisRepository {
	return &JurnalTeknoisRepository{Log: log}
}

func (r *JurnalTeknoisRepository) FindAll(tx *gorm.DB) ([]entity.JurnalTeknois, error) {
	var result []entity.JurnalTeknois
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
