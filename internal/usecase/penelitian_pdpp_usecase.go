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

func NewPenelitianPDPPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PenelitianPDPPRepository *repository.PenelitianPDPPRepository,
) *PenelitianPDPPUseCase {
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

	PenelitianPDPP := &entity.PenelitianPDPP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PenelitianPDPPRepository.Create(tx, PenelitianPDPP); err != nil {
		c.Log.WithError(err).Error("failed to create penelitian pdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PenelitianPDPPToResponse(PenelitianPDPP), nil
}

func (c *PenelitianPDPPUseCase) FindAll(ctx context.Context) ([]model.PenelitianPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianPDPP, err := c.PenelitianPDPPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting penelitian pdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting penelitian pdpp")
		return nil, err
	}

	responses := make([]model.PenelitianPDPPResponse, len(PenelitianPDPP))
	for i, penelitianPdpp := range PenelitianPDPP {
		responses[i] = *converter.PenelitianPDPPToResponse(&penelitianPdpp)
	}

	return responses, nil
}

func (c *PenelitianPDPPUseCase) Update(ctx context.Context, request *model.UpdatePenelitianPDPPRequest) (*model.PenelitianPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PenelitianPDPP := new(entity.PenelitianPDPP)
	if err := c.PenelitianPDPPRepository.FindById(tx, PenelitianPDPP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitian pdpp")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PenelitianPDPP.Title = request.Title
	PenelitianPDPP.Content = request.Content

	if err := c.PenelitianPDPPRepository.Update(tx, PenelitianPDPP); err != nil {
		c.Log.WithError(err).Error("error updating penelitian pdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating penelitian pdpp")
		return nil, err
	}

	return converter.PenelitianPDPPToResponse(PenelitianPDPP), nil
}

func (c *PenelitianPDPPUseCase) Delete(ctx context.Context, request *model.DeletePenelitianPDPPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	penelitianPDPP := new(entity.PenelitianPDPP)
	if err := c.PenelitianPDPPRepository.FindById(tx, penelitianPDPP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting penelitian pdpp")
		return err
	}

	if err := c.PenelitianPDPPRepository.Delete(tx, penelitianPDPP); err != nil {
		c.Log.WithError(err).Error("error deleting penelitian pdpp")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting penelitian pdpp")
		return err
	}

	return nil
}
