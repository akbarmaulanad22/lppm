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

type PKMHPPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMHPPRepository *repository.PKMHPPRepository
}

func NewPKMHPPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMHPPRepository *repository.PKMHPPRepository,
) *PKMHPPUseCase {
	return &PKMHPPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMHPPRepository: PKMHPPRepository,
	}
}

func (c *PKMHPPUseCase) Create(ctx context.Context, request *model.CreatePKMHPPRequest) (*model.PKMHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMHPP := &entity.PKMHPP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMHPPRepository.Create(tx, PKMHPP); err != nil {
		c.Log.WithError(err).Error("failed to create pkm hpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMHPPToResponse(PKMHPP), nil
}

func (c *PKMHPPUseCase) FindAll(ctx context.Context) ([]model.PKMHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMHPP, err := c.PKMHPPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm hpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm hpp")
		return nil, err
	}

	responses := make([]model.PKMHPPResponse, len(PKMHPP))
	for i, visiMisi := range PKMHPP {
		responses[i] = *converter.PKMHPPToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *PKMHPPUseCase) Update(ctx context.Context, request *model.UpdatePKMHPPRequest) (*model.PKMHPPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMHPP := new(entity.PKMHPP)
	if err := c.PKMHPPRepository.FindById(tx, PKMHPP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm hpp")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMHPP.Title = request.Title
	PKMHPP.Content = request.Content

	if err := c.PKMHPPRepository.Update(tx, PKMHPP); err != nil {
		c.Log.WithError(err).Error("error updating pkm hpp")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm hpp")
		return nil, err
	}

	return converter.PKMHPPToResponse(PKMHPP), nil
}

func (c *PKMHPPUseCase) Delete(ctx context.Context, request *model.DeletePKMHPPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	contact := new(entity.PKMHPP)
	if err := c.PKMHPPRepository.FindById(tx, contact, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	if err := c.PKMHPPRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return nil
}
