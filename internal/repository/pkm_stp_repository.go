package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PKMSTPRepository struct {
	Repository[entity.PKMSTP]
	Log *logrus.Logger
}

func NewPKMSTPRepository(log *logrus.Logger) *PKMSTPRepository {

	return &PKMSTPRepository{Log: log}

}

func (r *PKMSTPRepository) FindAll(tx *gorm.DB) ([]entity.PKMSTP, error) {
	var PKMSTP []entity.PKMSTP
	if err := tx.Find(&PKMSTP).Error; err != nil {
		return nil, err
	}
	return PKMSTP, nil
}
