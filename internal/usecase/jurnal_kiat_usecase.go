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

type JurnalKIATUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	JurnalKIATRepository *repository.JurnalKIATRepository
}

func NewJurnalKIATUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalKIATRepository *repository.JurnalKIATRepository) *JurnalKIATUseCase {
	return &JurnalKIATUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		JurnalKIATRepository: jurnalKIATRepository,
	}
}

func (c *JurnalKIATUseCase) Create(ctx context.Context, request *model.CreateJurnalKIATRequest) (*model.JurnalKIATResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalKIAT{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalKIATRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalKIAT")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalKIATToResponse(data), nil
}

func (c *JurnalKIATUseCase) FindAll(ctx context.Context) ([]model.JurnalKIATResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalKIATRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalKIAT")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalKIAT")
		return nil, err
	}

	responses := make([]model.JurnalKIATResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalKIATToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalKIATUseCase) Update(ctx context.Context, request *model.UpdateJurnalKIATRequest) (*model.JurnalKIATResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalKIAT)
	if err := c.JurnalKIATRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalKIAT")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalKIATRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalKIAT")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalKIAT")
		return nil, err
	}

	return converter.JurnalKIATToResponse(data), nil
}

func (c *JurnalKIATUseCase) Delete(ctx context.Context, request *model.DeleteJurnalKIATRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalKIAT)
	if err := c.JurnalKIATRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalKIAT")
		return err
	}

	if err := c.JurnalKIATRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalKIAT")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalKIAT")
		return err
	}

	return nil
}