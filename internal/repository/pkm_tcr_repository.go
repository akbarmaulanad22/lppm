package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMTCRRepository struct {
	Repository[entity.PKMTCR]
	Log *logrus.Logger
}

func NewPKMTCRRepository(log *logrus.Logger) *PKMTCRRepository {

	return &PKMTCRRepository{Log: log}

}

func (r *PKMTCRRepository) FindAll(tx *gorm.DB) ([]entity.PKMTCR, error) {
	var PKMTCR []entity.PKMTCR
	if err := tx.Find(&PKMTCR).Error; err != nil {
		return nil, err
	}
	return PKMTCR, nil
}
