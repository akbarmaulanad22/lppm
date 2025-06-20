package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HKIDosenRepository struct {
	Repository[entity.HKIDosen]
	Log *logrus.Logger
}

func NewHKIDosenRepository(log *logrus.Logger) *HKIDosenRepository {

	return &HKIDosenRepository{Log: log}

}

func (r *HKIDosenRepository) FindAll(tx *gorm.DB) ([]entity.HKIDosen, error) {
	var HKIDosen []entity.HKIDosen
	if err := tx.Find(&HKIDosen).Error; err != nil {
		return nil, err
	}
	return HKIDosen, nil
}
