package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMRDRPRepository struct {
	Repository[entity.PKMRDRP]
	Log *logrus.Logger
}

func NewPKMRDRPRepository(log *logrus.Logger) *PKMRDRPRepository {

	return &PKMRDRPRepository{Log: log}

}

func (r *PKMRDRPRepository) FindAll(tx *gorm.DB) ([]entity.PKMRDRP, error) {
	var PKMRDRP []entity.PKMRDRP
	if err := tx.Find(&PKMRDRP).Error; err != nil {
		return nil, err
	}
	return PKMRDRP, nil
}
