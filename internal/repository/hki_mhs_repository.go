package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HKIMHSRepository struct {
	Repository[entity.HKIMHS]
	Log *logrus.Logger
}

func NewHKIMHSRepository(log *logrus.Logger) *HKIMHSRepository {

	return &HKIMHSRepository{Log: log}

}

func (r *HKIMHSRepository) FindAll(tx *gorm.DB) ([]entity.HKIMHS, error) {
	var HKIMHS []entity.HKIMHS
	if err := tx.Find(&HKIMHS).Error; err != nil {
		return nil, err
	}
	return HKIMHS, nil
}
