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

type JurnalJSUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	JurnalJSRepository *repository.JurnalJSRepository
}

func NewJurnalJSUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalJSRepository *repository.JurnalJSRepository) *JurnalJSUseCase {
	return &JurnalJSUseCase{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		JurnalJSRepository: jurnalJSRepository,
	}
}

func (c *JurnalJSUseCase) Create(ctx context.Context, request *model.CreateJurnalJSRequest) (*model.JurnalJSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalJS{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalJSRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalJS")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalJSToResponse(data), nil
}

func (c *JurnalJSUseCase) FindAll(ctx context.Context) ([]model.JurnalJSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalJSRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalJS")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalJS")
		return nil, err
	}

	responses := make([]model.JurnalJSResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalJSToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalJSUseCase) Update(ctx context.Context, request *model.UpdateJurnalJSRequest) (*model.JurnalJSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalJS)
	if err := c.JurnalJSRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalJS")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalJSRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalJS")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalJS")
		return nil, err
	}

	return converter.JurnalJSToResponse(data), nil
}

func (c *JurnalJSUseCase) Delete(ctx context.Context, request *model.DeleteJurnalJSRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalJS)
	if err := c.JurnalJSRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalJS")
		return err
	}

	if err := c.JurnalJSRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalJS")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalJS")
		return err
	}

	return nil
} 