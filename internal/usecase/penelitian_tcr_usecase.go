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

type PenelitianTCRUseCase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	Validate                *validator.Validate
	PenelitianTCRRepository *repository.PenelitianTCRRepository
}

func NewPenelitianTCRUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PenelitianTCRRepository *repository.PenelitianTCRRepository,
) *PenelitianTCRUseCase {
	return &PenelitianTCRUseCase{
		DB:                      db,
		Log:                     logger,
		Validate:                validate,
		PenelitianTCRRepository: PenelitianTCRRepository,
	}
}

func (c *PenelitianTCRUseCase) Create(ctx context.Context, request *model.CreatePenelitianTCRRequest) (*model.PenelitianTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PenelitianTCR := &entity.PenelitianTCR{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianTCRRepository.Create(tx, PenelitianTCR); err != nil {
		c.Log.WithError(err).Error("failed to create penelitian tcr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianTCRToResponse(PenelitianTCR), nil
}

func (c *PenelitianTCRUseCase) FindAll(ctx context.Context) ([]model.PenelitianTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianTCR, err := c.PenelitianTCRRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting penelitian tcr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting penelitian tcr")
		return nil, err
	}

	responses := make([]model.PenelitianTCRResponse, len(PenelitianTCR))
	for i, visiMisi := range PenelitianTCR {
		responses[i] = *converter.PenelitianTCRToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *PenelitianTCRUseCase) Update(ctx context.Context, request *model.UpdatePenelitianTCRRequest) (*model.PenelitianTCRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianTCR := new(entity.PenelitianTCR)
	if err := c.PenelitianTCRRepository.FindById(tx, PenelitianTCR, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitian tcr")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PenelitianTCR.Title = request.Title
	PenelitianTCR.Content = request.Content

	if err := c.PenelitianTCRRepository.Update(tx, PenelitianTCR); err != nil {
		c.Log.WithError(err).Error("error updating penelitian tcr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating penelitian tcr")
		return nil, err
	}

	return converter.PenelitianTCRToResponse(PenelitianTCR), nil
}

func (c *PenelitianTCRUseCase) Delete(ctx context.Context, request *model.DeletePenelitianTCRRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	penelitianTCR := new(entity.PenelitianTCR)
	if err := c.PenelitianTCRRepository.FindById(tx, penelitianTCR, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitian tcr")
		return err
	}

	if err := c.PenelitianTCRRepository.Delete(tx, penelitianTCR); err != nil {
		c.Log.WithError(err).Error("error deleting penelitian tcr")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting penelitian tcr")
		return err
	}

	return nil
}
