package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMLPRepository struct {
	Repository[entity.PKMLP]
	Log *logrus.Logger
}

func NewPKMLPRepository(log *logrus.Logger) *PKMLPRepository {

	return &PKMLPRepository{Log: log}

}

func (r *PKMLPRepository) FindAll(tx *gorm.DB) ([]entity.PKMLP, error) {
	var PKMLP []entity.PKMLP
	if err := tx.Find(&PKMLP).Error; err != nil {
		return nil, err
	}
	return PKMLP, nil
}
