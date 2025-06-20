package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalJSRepository struct {
	Repository[entity.JurnalJS]
	Log *logrus.Logger
}

func NewJurnalJSRepository(log *logrus.Logger) *JurnalJSRepository {
	return &JurnalJSRepository{Log: log}
}

func (r *JurnalJSRepository) FindAll(tx *gorm.DB) ([]entity.JurnalJS, error) {
	var result []entity.JurnalJS
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
} 