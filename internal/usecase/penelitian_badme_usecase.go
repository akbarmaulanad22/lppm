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

type PenelitianBADMEUseCase struct {
	DB                        *gorm.DB
	Log                       *logrus.Logger
	Validate                  *validator.Validate
	PenelitianBADMERepository *repository.PenelitianBADMERepository
}

func NewPenelitianBADMEUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianBADMERepository *repository.PenelitianBADMERepository) *PenelitianBADMEUseCase {
	return &PenelitianBADMEUseCase{
		DB:                        db,
		Log:                       logger,
		Validate:                  validate,
		PenelitianBADMERepository: PenelitianBADMERepository,
	}
}

func (c *PenelitianBADMEUseCase) Create(ctx context.Context, request *model.CreatePenelitianBADMERequest) (*model.PenelitianBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianBADME{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianBADMERepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianBADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianBADMEToResponse(data), nil
}

func (c *PenelitianBADMEUseCase) FindAll(ctx context.Context) ([]model.PenelitianBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianBADMERepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianBADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianBADME")
		return nil, err
	}

	responses := make([]model.PenelitianBADMEResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianBADMEToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianBADMEUseCase) Update(ctx context.Context, request *model.UpdatePenelitianBADMERequest) (*model.PenelitianBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianBADME)
	if err := c.PenelitianBADMERepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianBADME")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianBADMERepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianBADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianBADME")
		return nil, err
	}

	return converter.PenelitianBADMEToResponse(data), nil
}

func (c *PenelitianBADMEUseCase) Delete(ctx context.Context, request *model.DeletePenelitianBADMERequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianBADME)
	if err := c.PenelitianBADMERepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianBADME")
		return err
	}

	if err := c.PenelitianBADMERepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianBADME")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianBADME")
		return err
	}

	return nil
}
