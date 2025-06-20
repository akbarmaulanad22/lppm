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

type PenelitianPDPPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PenelitianPDPPRepository *repository.PenelitianPDPPRepository
}

func NewPenelitianPDPPUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, PenelitianPDPPRepository *repository.PenelitianPDPPRepository) *PenelitianPDPPUseCase {
	return &PenelitianPDPPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PenelitianPDPPRepository: PenelitianPDPPRepository,
	}
}

func (c *PenelitianPDPPUseCase) Create(ctx context.Context, request *model.CreatePenelitianPDPPRequest) (*model.PenelitianPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	data := &entity.PenelitianPDPP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianPDPPRepository.Create(tx, data); err != nil {
		c.Log.WithError(err).Error("failed to create PenelitianPDPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianPDPPToResponse(data), nil
}

func (c *PenelitianPDPPUseCase) FindAll(ctx context.Context) ([]model.PenelitianPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	result, err := c.PenelitianPDPPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting PenelitianPDPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting PenelitianPDPP")
		return nil, err
	}

	responses := make([]model.PenelitianPDPPResponse, len(result))
	for i, v := range result {
		responses[i] = *converter.PenelitianPDPPToResponse(&v)
	}

	return responses, nil
}

func (c *PenelitianPDPPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianPDPPRequest) (*model.PenelitianPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data := new(entity.PenelitianPDPP)
	if err := c.PenelitianPDPPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianPDPP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	data.Title = request.Title
	data.Content = request.Content

	if err := c.PenelitianPDPPRepository.Update(tx, data); err != nil {
		c.Log.WithError(err).Error("error updating PenelitianPDPP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating PenelitianPDPP")
		return nil, err
	}

	return converter.PenelitianPDPPToResponse(data), nil
}

func (c *PenelitianPDPPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianPDPPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	data := new(entity.PenelitianPDPP)
	if err := c.PenelitianPDPPRepository.FindById(tx, data, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting PenelitianPDPP")
		return err
	}

	if err := c.PenelitianPDPPRepository.Delete(tx, data); err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianPDPP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting PenelitianPDPP")
		return err
	}

	return nil
}
