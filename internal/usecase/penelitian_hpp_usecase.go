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

type PenelitianHPPUseCase struct {
	DB *gorm.DB
	Log *logrus.Logger
	Validate *validator.Validate
	PenelitianHPPRepository *repository.PenelitianHPPRepository
}

func NewPenelitianHPPUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianHPPRepository *repository.PenelitianHPPRepository) *PenelitianHPPUseCase {
	return &PenelitianHPPUseCase{
		DB: db,
		Log: logger,
		Validate: validate,
		PenelitianHPPRepository: PenelitianHPPRepository,
	}
}

func (c *PenelitianHPPUseCase) Create(ctx context.Context, request *model.CreatePenelitianHPPRequest) (*model.PenelitianHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianHPP{
		Title: request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianHPPRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianHPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianHPPToResponse(data), nil
}

func (c *PenelitianHPPUseCase) FindAll(ctx context.Context) ([]model.PenelitianHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianHPPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianHPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianHPP")
		return nil, err
	}

	responses := make([]model.PenelitianHPPResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianHPPToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianHPPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianHPPRequest) (*model.PenelitianHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianHPP)
	if err := c.PenelitianHPPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianHPP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianHPPRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianHPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianHPP")
		return nil, err
	}

	return converter.PenelitianHPPToResponse(data), nil
}

func (c *PenelitianHPPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianHPPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianHPP)
	if err := c.PenelitianHPPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianHPP")
		return err
	}

	if err := c.PenelitianHPPRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianHPP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianHPP")
		return err
	}

	return nil
} 