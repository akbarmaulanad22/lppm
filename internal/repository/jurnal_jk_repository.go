package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalJKRepository struct {
	Repository[entity.JurnalJK]
	Log *logrus.Logger
}

func NewJurnalJKRepository(log *logrus.Logger) *JurnalJKRepository {
	return &JurnalJKRepository{Log: log}
}

func (r *JurnalJKRepository) FindAll(tx *gorm.DB) ([]entity.JurnalJK, error) {
	var result []entity.JurnalJK
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
