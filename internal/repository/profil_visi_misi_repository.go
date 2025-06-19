package repository

import (
	"tugasakhir/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfilVisiMisiRepository struct {
	Repository[entity.ProfilVisiMisi]
	Log *logrus.Logger
}

func NewProfilVisiMisiRepository(log *logrus.Logger) *ProfilVisiMisiRepository {

	return &ProfilVisiMisiRepository{Log: log}

}

func (r *ProfilVisiMisiRepository) FindAll(tx *gorm.DB) ([]entity.ProfilVisiMisi, error) {
	var profilVisiMisi []entity.ProfilVisiMisi
	if err := tx.Find(&profilVisiMisi).Error; err != nil {
		return nil, err
	}
	return profilVisiMisi, nil
}
