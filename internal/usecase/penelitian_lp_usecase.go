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

type PenelitianLPUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	PenelitianLPRepository *repository.PenelitianLPRepository
}

func NewPenelitianLPUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianLPRepository *repository.PenelitianLPRepository) *PenelitianLPUseCase {
	return &PenelitianLPUseCase{
		DB:                     db,
		Log:                    logger,
		Validate:               validate,
		PenelitianLPRepository: PenelitianLPRepository,
	}
}

func (c *PenelitianLPUseCase) Create(ctx context.Context, request *model.CreatePenelitianLPRequest) (*model.PenelitianLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianLP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianLPRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianLP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianLPToResponse(data), nil
}

func (c *PenelitianLPUseCase) FindAll(ctx context.Context) ([]model.PenelitianLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianLPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianLP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianLP")
		return nil, err
	}

	responses := make([]model.PenelitianLPResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianLPToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianLPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianLPRequest) (*model.PenelitianLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianLP)
	if err := c.PenelitianLPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianLP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianLPRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianLP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianLP")
		return nil, err
	}

	return converter.PenelitianLPToResponse(data), nil
}

func (c *PenelitianLPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianLPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianLP)
	if err := c.PenelitianLPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianLP")
		return err
	}

	if err := c.PenelitianLPRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianLP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianLP")
		return err
	}

	return nil
}
