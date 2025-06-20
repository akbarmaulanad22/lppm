package usecase

import (
	"context"
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
	"tugasakhir/internal/model/converter"
	"tugasakhir/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfilVisiMisiUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	ProfilVisiMisiRepository *repository.ProfilVisiMisiRepository
}

func NewProfilVisiMisiUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	profilVisiMisiRepository *repository.ProfilVisiMisiRepository,
) *ProfilVisiMisiUseCase {
	return &ProfilVisiMisiUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		ProfilVisiMisiRepository: profilVisiMisiRepository,
	}
}

func (c *ProfilVisiMisiUseCase) Create(ctx context.Context, request *model.CreateProfilVisiMisiRequest) (*model.ProfilVisiMisiResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	profilVisiMisi := &entity.ProfilVisiMisi{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.ProfilVisiMisiRepository.Create(tx, profilVisiMisi); err != nil {
		c.Log.WithError(err).Error("failed to create profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.ProfilVisiMisiToResponse(profilVisiMisi), nil
}

func (c *ProfilVisiMisiUseCase) FindAll(ctx context.Context) ([]model.ProfilVisiMisiResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	profilVisiMisi, err := c.ProfilVisiMisiRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	responses := make([]model.ProfilVisiMisiResponse, len(profilVisiMisi))
	for i, visiMisi := range profilVisiMisi {
		responses[i] = *converter.ProfilVisiMisiToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *ProfilVisiMisiUseCase) Update(ctx context.Context, request *model.UpdateProfilVisiMisiRequest) (*model.ProfilVisiMisiResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	profilVisiMisi := new(entity.ProfilVisiMisi)
	if err := c.ProfilVisiMisiRepository.FindById(tx, profilVisiMisi, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	profilVisiMisi.Title = request.Title
	profilVisiMisi.Content = request.Content

	if err := c.ProfilVisiMisiRepository.Update(tx, profilVisiMisi); err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	return converter.ProfilVisiMisiToResponse(profilVisiMisi), nil
}

func (c *ProfilVisiMisiUseCase) Delete(ctx context.Context, request *model.DeleteProfilVisiMisiRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	profilVISIMISI := new(entity.ProfilVisiMisi)
	if err := c.ProfilVisiMisiRepository.FindById(tx, profilVISIMISI, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profilVISIMISI")
		return err
	}

	if err := c.ProfilVisiMisiRepository.Delete(tx, profilVISIMISI); err != nil {
		c.Log.WithError(err).Error("error deleting profilVISIMISI")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting profilVISIMISI")
		return err
	}

	return nil
}
