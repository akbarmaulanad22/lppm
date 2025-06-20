package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMSKRRepository struct {
	Repository[entity.PKMSKR]
	Log *logrus.Logger
}

func NewPKMSKRRepository(log *logrus.Logger) *PKMSKRRepository {

	return &PKMSKRRepository{Log: log}

}

func (r *PKMSKRRepository) FindAll(tx *gorm.DB) ([]entity.PKMSKR, error) {
	var PKMSKR []entity.PKMSKR
	if err := tx.Find(&PKMSKR).Error; err != nil {
		return nil, err
	}
	return PKMSKR, nil
}
