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

type JurnalTAJBUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	JurnalTAJBRepository *repository.JurnalTAJBRepository
}

func NewJurnalTAJBUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalTAJBRepository *repository.JurnalTAJBRepository) *JurnalTAJBUseCase {
	return &JurnalTAJBUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		JurnalTAJBRepository: jurnalTAJBRepository,
	}
}

func (c *JurnalTAJBUseCase) Create(ctx context.Context, request *model.CreateJurnalTAJBRequest) (*model.JurnalTAJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalTAJB{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalTAJBRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalTAJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalTAJBToResponse(data), nil
}

func (c *JurnalTAJBUseCase) FindAll(ctx context.Context) ([]model.JurnalTAJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalTAJBRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalTAJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalTAJB")
		return nil, err
	}

	responses := make([]model.JurnalTAJBResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalTAJBToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalTAJBUseCase) Update(ctx context.Context, request *model.UpdateJurnalTAJBRequest) (*model.JurnalTAJBResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalTAJB)
	if err := c.JurnalTAJBRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTAJB")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalTAJBRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalTAJB")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalTAJB")
		return nil, err
	}

	return converter.JurnalTAJBToResponse(data), nil
}

func (c *JurnalTAJBUseCase) Delete(ctx context.Context, request *model.DeleteJurnalTAJBRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalTAJB)
	if err := c.JurnalTAJBRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalTAJB")
		return err
	}

	if err := c.JurnalTAJBRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTAJB")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalTAJB")
		return err
	}

	return nil
}
