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

type PKMPDPPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMPDPPRepository *repository.PKMPDPPRepository
}

func NewPKMPDPPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMPDPPRepository *repository.PKMPDPPRepository,
) *PKMPDPPUseCase {
	return &PKMPDPPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMPDPPRepository: PKMPDPPRepository,
	}
}

func (c *PKMPDPPUseCase) Create(ctx context.Context, request *model.CreatePKMPDPPRequest) (*model.PKMPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMPDPP := &entity.PKMPDPP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMPDPPRepository.Create(tx, PKMPDPP); err != nil {
		c.Log.WithError(err).Error("failed to create pkm hdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMPDPPToResponse(PKMPDPP), nil
}

func (c *PKMPDPPUseCase) FindAll(ctx context.Context) ([]model.PKMPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMPDPP, err := c.PKMPDPPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm hdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm hdpp")
		return nil, err
	}

	responses := make([]model.PKMPDPPResponse, len(PKMPDPP))
	for i, visiMisi := range PKMPDPP {
		responses[i] = *converter.PKMPDPPToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *PKMPDPPUseCase) Update(ctx context.Context, request *model.UpdatePKMPDPPRequest) (*model.PKMPDPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMPDPP := new(entity.PKMPDPP)
	if err := c.PKMPDPPRepository.FindById(tx, PKMPDPP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm hdpp")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMPDPP.Title = request.Title
	PKMPDPP.Content = request.Content

	if err := c.PKMPDPPRepository.Update(tx, PKMPDPP); err != nil {
		c.Log.WithError(err).Error("error updating pkm hdpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm hdpp")
		return nil, err
	}

	return converter.PKMPDPPToResponse(PKMPDPP), nil
}

func (c *PKMPDPPUseCase) Delete(ctx context.Context, request *model.DeletePKMPDPPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	contact := new(entity.PKMPDPP)
	if err := c.PKMPDPPRepository.FindById(tx, contact, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	if err := c.PKMPDPPRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return nil
}
