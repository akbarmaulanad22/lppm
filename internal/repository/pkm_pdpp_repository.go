package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMPDPPRepository struct {
	Repository[entity.PKMPDPP]
	Log *logrus.Logger
}

func NewPKMPDPPRepository(log *logrus.Logger) *PKMPDPPRepository {

	return &PKMPDPPRepository{Log: log}

}

func (r *PKMPDPPRepository) FindAll(tx *gorm.DB) ([]entity.PKMPDPP, error) {
	var PKMPDPP []entity.PKMPDPP
	if err := tx.Find(&PKMPDPP).Error; err != nil {
		return nil, err
	}
	return PKMPDPP, nil
}
