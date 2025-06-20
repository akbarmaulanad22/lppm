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

type JurnalJKUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	JurnalJKRepository *repository.JurnalJKRepository
}

func NewJurnalJKUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, jurnalJKRepository *repository.JurnalJKRepository) *JurnalJKUseCase {
	return &JurnalJKUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		JurnalJKRepository: jurnalJKRepository,
	}
}

func (c *JurnalJKUseCase) Create(ctx context.Context, request *model.CreateJurnalJKRequest) (*model.JurnalJKResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.JurnalJK{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.JurnalJKRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create JurnalJK")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.JurnalJKToResponse(data), nil
}

func (c *JurnalJKUseCase) FindAll(ctx context.Context) ([]model.JurnalJKResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.JurnalJKRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting JurnalJK")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting JurnalJK")
		return nil, err
	}

	responses := make([]model.JurnalJKResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.JurnalJKToResponse(&v)
	}

	return responses, nil
}

func (c *JurnalJKUseCase) Update(ctx context.Context, request *model.UpdateJurnalJKRequest) (*model.JurnalJKResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.JurnalJK)
	if err := c.JurnalJKRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalJK")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.JurnalJKRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating JurnalJK")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating JurnalJK")
		return nil, err
	}

	return converter.JurnalJKToResponse(data), nil
}

func (c *JurnalJKUseCase) Delete(ctx context.Context, request *model.DeleteJurnalJKRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.JurnalJK)
	if err := c.JurnalJKRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting JurnalJK")
		return err
	}

	if err := c.JurnalJKRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting JurnalJK")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting JurnalJK")
		return err
	}

	return nil
}
