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

type PKMTCRUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMTCRRepository *repository.PKMTCRRepository
}

func NewPKMTCRUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMTCRRepository *repository.PKMTCRRepository,
) *PKMTCRUseCase {
	return &PKMTCRUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMTCRRepository: PKMTCRRepository,
	}
}

func (c *PKMTCRUseCase) Create(ctx context.Context, request *model.CreatePKMTCRRequest) (*model.PKMTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMTCR := &entity.PKMTCR{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMTCRRepository.Create(tx, PKMTCR); err != nil {
		c.Log.WithError(err).Error("failed to create profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMTCRToResponse(PKMTCR), nil
}

func (c *PKMTCRUseCase) FindAll(ctx context.Context) ([]model.PKMTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMTCR, err := c.PKMTCRRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	responses := make([]model.PKMTCRResponse, len(PKMTCR))
	for i, visiMisi := range PKMTCR {
		responses[i] = *converter.PKMTCRToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *PKMTCRUseCase) Update(ctx context.Context, request *model.UpdatePKMTCRRequest) (*model.PKMTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMTCR := new(entity.PKMTCR)
	if err := c.PKMTCRRepository.FindById(tx, PKMTCR, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMTCR.Title = request.Title
	PKMTCR.Content = request.Content

	if err := c.PKMTCRRepository.Update(tx, PKMTCR); err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	return converter.PKMTCRToResponse(PKMTCR), nil
}

func (c *PKMTCRUseCase) Delete(ctx context.Context, request *model.DeletePKMTCRRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	contact := new(entity.PKMTCR)
	if err := c.PKMTCRRepository.FindById(tx, contact, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	if err := c.PKMTCRRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return nil
}
