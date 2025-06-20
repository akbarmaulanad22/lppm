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

type PenelitianSTPUseCase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	Validate                *validator.Validate
	PenelitianSTPRepository *repository.PenelitianSTPRepository
}

func NewPenelitianSTPUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianSTPRepository *repository.PenelitianSTPRepository) *PenelitianSTPUseCase {
	return &PenelitianSTPUseCase{
		DB:                      db,
		Log:                     logger,
		Validate:                validate,
		PenelitianSTPRepository: PenelitianSTPRepository,
	}
}

func (c *PenelitianSTPUseCase) Create(ctx context.Context, request *model.CreatePenelitianSTPRequest) (*model.PenelitianSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianSTP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianSTPRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianSTP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianSTPToResponse(data), nil
}

func (c *PenelitianSTPUseCase) FindAll(ctx context.Context) ([]model.PenelitianSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianSTPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSTP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSTP")
		return nil, err
	}

	responses := make([]model.PenelitianSTPResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianSTPToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianSTPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianSTPRequest) (*model.PenelitianSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianSTP)
	if err := c.PenelitianSTPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSTP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianSTPRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianSTP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianSTP")
		return nil, err
	}

	return converter.PenelitianSTPToResponse(data), nil
}

func (c *PenelitianSTPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianSTPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianSTP)
	if err := c.PenelitianSTPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianSTP")
		return err
	}

	if err := c.PenelitianSTPRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianSTP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianSTP")
		return err
	}

	return nil
}
