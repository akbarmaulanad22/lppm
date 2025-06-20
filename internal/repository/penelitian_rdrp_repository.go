package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PenelitianRDRPRepository struct {
	Repository[entity.PenelitianRDRP]
	Log *logrus.Logger
}

func NewPenelitianRDRPRepository(log *logrus.Logger) *PenelitianRDRPRepository {

	return &PenelitianRDRPRepository{Log: log}

}

func (r *PenelitianRDRPRepository) FindAll(tx *gorm.DB) ([]entity.PenelitianRDRP, error) {
	var PenelitianRDRP []entity.PenelitianRDRP
	if err := tx.Find(&PenelitianRDRP).Error; err != nil {
		return nil, err
	}
	return PenelitianRDRP, nil
}
