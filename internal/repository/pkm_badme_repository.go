package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMBADMERepository struct {
	Repository[entity.PKMBADME]
	Log *logrus.Logger
}

func NewPKMBADMERepository(log *logrus.Logger) *PKMBADMERepository {

	return &PKMBADMERepository{Log: log}

}

func (r *PKMBADMERepository) FindAll(tx *gorm.DB) ([]entity.PKMBADME, error) {
	var PKMBADME []entity.PKMBADME
	if err := tx.Find(&PKMBADME).Error; err != nil {
		return nil, err
	}
	return PKMBADME, nil
}
