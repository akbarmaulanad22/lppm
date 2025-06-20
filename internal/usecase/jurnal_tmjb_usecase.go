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

type JurnalTMJBUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	JurnalTMJBRepository *repository.JurnalTMJBRepository
}

func NewJurnalTMJBUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalTMJBRepository *repository.JurnalTMJBRepository) *JurnalTMJBUseCase {
	return &JurnalTMJBUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		JurnalTMJBRepository: jurnalTMJBRepository,
	}
}

func (c *JurnalTMJBUseCase) Create(ctx context.Context, request *model.CreateJurnalTMJBRequest) (*model.JurnalTMJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalTMJB{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalTMJBRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalTMJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalTMJBToResponse(data), nil
}

func (c *JurnalTMJBUseCase) FindAll(ctx context.Context) ([]model.JurnalTMJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalTMJBRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalTMJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalTMJB")
		return nil, err
	}

	responses := make([]model.JurnalTMJBResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalTMJBToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalTMJBUseCase) Update(ctx context.Context, request *model.UpdateJurnalTMJBRequest) (*model.JurnalTMJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalTMJB)
	if err := c.JurnalTMJBRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTMJB")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalTMJBRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalTMJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalTMJB")
		return nil, err
	}

	return converter.JurnalTMJBToResponse(data), nil
}

func (c *JurnalTMJBUseCase) Delete(ctx context.Context, request *model.DeleteJurnalTMJBRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalTMJB)
	if err := c.JurnalTMJBRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTMJB")
		return err
	}

	if err := c.JurnalTMJBRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTMJB")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTMJB")
		return err
	}

	return nil
}
