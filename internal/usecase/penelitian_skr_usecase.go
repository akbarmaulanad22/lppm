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

type PenelitianSKRUseCase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	Validate                *validator.Validate
	PenelitianSKRRepository *repository.PenelitianSKRRepository
}

func NewPenelitianSKRUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianSKRRepository *repository.PenelitianSKRRepository) *PenelitianSKRUseCase {
	return &PenelitianSKRUseCase{
		DB:                      db,
		Log:                     logger,
		Validate:                validate,
		PenelitianSKRRepository: PenelitianSKRRepository,
	}
}

func (c *PenelitianSKRUseCase) Create(ctx context.Context, request *model.CreatePenelitianSKRRequest) (*model.PenelitianSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianSKR{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianSKRRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianSKR")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianSKRToResponse(data), nil
}

func (c *PenelitianSKRUseCase) FindAll(ctx context.Context) ([]model.PenelitianSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianSKRRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSKR")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSKR")
		return nil, err
	}

	responses := make([]model.PenelitianSKRResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianSKRToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianSKRUseCase) Update(ctx context.Context, request *model.UpdatePenelitianSKRRequest) (*model.PenelitianSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianSKR)
	if err := c.PenelitianSKRRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSKR")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianSKRRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianSKR")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianSKR")
		return nil, err
	}

	return converter.PenelitianSKRToResponse(data), nil
}

func (c *PenelitianSKRUseCase) Delete(ctx context.Context, request *model.DeletePenelitianSKRRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianSKR)
	if err := c.PenelitianSKRRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSKR")
		return err
	}

	if err := c.PenelitianSKRRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianSKR")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianSKR")
		return err
	}

	return nil
}
