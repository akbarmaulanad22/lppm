package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JurnalTMJBRepository struct {
	Repository[entity.JurnalTMJB]
	Log *logrus.Logger
}

func NewJurnalTMJBRepository(log *logrus.Logger) *JurnalTMJBRepository {
	return &JurnalTMJBRepository{Log: log}
}

func (r *JurnalTMJBRepository) FindAll(tx *gorm.DB) ([]entity.JurnalTMJB, error) {
	var result []entity.JurnalTMJB
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
