package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfilSODTRepository struct {
	Repository[entity.ProfilSODT]
	Log *logrus.Logger
}

func NewProfilSODTRepository(log *logrus.Logger) *ProfilSODTRepository {

	return &ProfilSODTRepository{Log: log}

}

func (r *ProfilSODTRepository) FindAll(tx *gorm.DB) ([]entity.ProfilSODT, error) {
	var profilSODT []entity.ProfilSODT
	if err := tx.Find(&profilSODT).Error; err != nil {
		return nil, err
	}
	return profilSODT, nil
}
