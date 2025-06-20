package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMHPPRepository struct {
	Repository[entity.PKMHPP]
	Log *logrus.Logger
}

func NewPKMHPPRepository(log *logrus.Logger) *PKMHPPRepository {

	return &PKMHPPRepository{Log: log}

}

func (r *PKMHPPRepository) FindAll(tx *gorm.DB) ([]entity.PKMHPP, error) {
	var PKMHPP []entity.PKMHPP
	if err := tx.Find(&PKMHPP).Error; err != nil {
		return nil, err
	}
	return PKMHPP, nil
}
