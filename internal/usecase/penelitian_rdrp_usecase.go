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

type PenelitianRDRPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PenelitianRDRPRepository *repository.PenelitianRDRPRepository
}

func NewPenelitianRDRPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PenelitianRDRPRepository *repository.PenelitianRDRPRepository,
) *PenelitianRDRPUseCase {
	return &PenelitianRDRPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PenelitianRDRPRepository: PenelitianRDRPRepository,
	}
}

func (c *PenelitianRDRPUseCase) Create(ctx context.Context, request *model.CreatePenelitianRDRPRequest) (*model.PenelitianRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PenelitianRDRP := &entity.PenelitianRDRP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianRDRPRepository.Create(tx, PenelitianRDRP); err != nil {
		c.Log.WithError(err).Error("failed to create penelitian rdrp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianRDRPToResponse(PenelitianRDRP), nil
}

func (c *PenelitianRDRPUseCase) FindAll(ctx context.Context) ([]model.PenelitianRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianRDRP, err := c.PenelitianRDRPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting penelitian rdrp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting penelitian rdrp")
		return nil, err
	}

	responses := make([]model.PenelitianRDRPResponse, len(PenelitianRDRP))
	for i, penelitianRdrp := range PenelitianRDRP {
		responses[i] = *converter.PenelitianRDRPToResponse(&penelitianRdrp)
	}

	return responses, nil
}

func (c *PenelitianRDRPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianRDRPRequest) (*model.PenelitianRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianRDRP := new(entity.PenelitianRDRP)
	if err := c.PenelitianRDRPRepository.FindById(tx, PenelitianRDRP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitian rdrp")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PenelitianRDRP.Title = request.Title
	PenelitianRDRP.Content = request.Content

	if err := c.PenelitianRDRPRepository.Update(tx, PenelitianRDRP); err != nil {
		c.Log.WithError(err).Error("error updating penelitian rdrp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating penelitian rdrp")
		return nil, err
	}

	return converter.PenelitianRDRPToResponse(PenelitianRDRP), nil
}

func (c *PenelitianRDRPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianRDRPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	penelitianRDRP := new(entity.PenelitianRDRP)
	if err := c.PenelitianRDRPRepository.FindById(tx, penelitianRDRP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitianRDRP")
		return err
	}

	if err := c.PenelitianRDRPRepository.Delete(tx, penelitianRDRP); err != nil {
		c.Log.WithError(err).Error("error deleting penelitianRDRP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting penelitianRDRP")
		return err
	}

	return nil
}
