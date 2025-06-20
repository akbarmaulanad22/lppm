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

type ProfilSODTUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	ProfilSODTRepository *repository.ProfilSODTRepository
}

func NewProfilSODTUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	profilSODTRepository *repository.ProfilSODTRepository,
) *ProfilSODTUseCase {
	return &ProfilSODTUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		ProfilSODTRepository: profilSODTRepository,
	}
}

func (c *ProfilSODTUseCase) Create(ctx context.Context, request *model.CreateProfilSODTRequest) (*model.ProfilSODTResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	profilSODT := &entity.ProfilSODT{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.ProfilSODTRepository.Create(tx, profilSODT); err != nil {
		c.Log.WithError(err).Error("failed to create profil sodt")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.ProfilSODTToResponse(profilSODT), nil
}

func (c *ProfilSODTUseCase) FindAll(ctx context.Context) ([]model.ProfilSODTResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	profilSODT, err := c.ProfilSODTRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting profil sodt")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting profil sodt")
		return nil, err
	}

	responses := make([]model.ProfilSODTResponse, len(profilSODT))
	for i, sodt := range profilSODT {
		responses[i] = *converter.ProfilSODTToResponse(&sodt)
	}

	return responses, nil
}

func (c *ProfilSODTUseCase) Update(ctx context.Context, request *model.UpdateProfilSODTRequest) (*model.ProfilSODTResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	profilSODT := new(entity.ProfilSODT)
	if err := c.ProfilSODTRepository.FindById(tx, profilSODT, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profil sodt")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	profilSODT.Title = request.Title
	profilSODT.Content = request.Content

	if err := c.ProfilSODTRepository.Update(tx, profilSODT); err != nil {
		c.Log.WithError(err).Error("error updating profil sodt")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating profil sodt")
		return nil, err
	}

	return converter.ProfilSODTToResponse(profilSODT), nil
}

func (c *ProfilSODTUseCase) Delete(ctx context.Context, request *model.DeleteProfilSODTRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	profilSODT := new(entity.ProfilSODT)
	if err := c.ProfilSODTRepository.FindById(tx, profilSODT, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profilSODT")
		return err
	}

	if err := c.ProfilSODTRepository.Delete(tx, profilSODT); err != nil {
		c.Log.WithError(err).Error("error deleting profilSODT")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting profilSODT")
		return err
	}

	return nil
}
