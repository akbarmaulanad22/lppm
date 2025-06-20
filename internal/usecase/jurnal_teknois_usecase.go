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

type JurnalTeknoisUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	JurnalTeknoisRepository *repository.JurnalTeknoisRepository
}

func NewJurnalTeknoisUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalTeknoisRepository *repository.JurnalTeknoisRepository) *JurnalTeknoisUseCase {
	return &JurnalTeknoisUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		JurnalTeknoisRepository: jurnalTeknoisRepository,
	}
}

func (c *JurnalTeknoisUseCase) Create(ctx context.Context, request *model.CreateJurnalTeknoisRequest) (*model.JurnalTeknoisResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalTeknois{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalTeknoisRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalTeknois")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalTeknoisToResponse(data), nil
}

func (c *JurnalTeknoisUseCase) FindAll(ctx context.Context) ([]model.JurnalTeknoisResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalTeknoisRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalTeknois")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalTeknois")
		return nil, err
	}

	responses := make([]model.JurnalTeknoisResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalTeknoisToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalTeknoisUseCase) Update(ctx context.Context, request *model.UpdateJurnalTeknoisRequest) (*model.JurnalTeknoisResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalTeknois)
	if err := c.JurnalTeknoisRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTeknois")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalTeknoisRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalTeknois")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalTeknois")
		return nil, err
	}

	return converter.JurnalTeknoisToResponse(data), nil
}

func (c *JurnalTeknoisUseCase) Delete(ctx context.Context, request *model.DeleteJurnalTeknoisRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalTeknois)
	if err := c.JurnalTeknoisRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTeknois")
		return err
	}

	if err := c.JurnalTeknoisRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTeknois")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTeknois")
		return err
	}

	return nil
} 