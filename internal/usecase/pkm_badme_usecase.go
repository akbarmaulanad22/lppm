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

type PKMBADMEUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMBADMERepository *repository.PKMBADMERepository
}

func NewPKMBADMEUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMBADMERepository *repository.PKMBADMERepository,
) *PKMBADMEUseCase {
	return &PKMBADMEUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMBADMERepository: PKMBADMERepository,
	}
}

func (c *PKMBADMEUseCase) Create(ctx context.Context, request *model.CreatePKMBADMERequest) (*model.PKMBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMBADME := &entity.PKMBADME{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMBADMERepository.Create(tx, PKMBADME); err != nil {
		c.Log.WithError(err).Error("failed to create pkm BADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMBADMEToResponse(PKMBADME), nil
}

func (c *PKMBADMEUseCase) FindAll(ctx context.Context) ([]model.PKMBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMBADME, err := c.PKMBADMERepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm BADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm BADME")
		return nil, err
	}

	responses := make([]model.PKMBADMEResponse, len(PKMBADME))
	for i, pkmBadme := range PKMBADME {
		responses[i] = *converter.PKMBADMEToResponse(&pkmBadme)
	}

	return responses, nil
}

func (c *PKMBADMEUseCase) Update(ctx context.Context, request *model.UpdatePKMBADMERequest) (*model.PKMBADMEResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMBADME := new(entity.PKMBADME)
	if err := c.PKMBADMERepository.FindById(tx, PKMBADME, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm BADME")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMBADME.Title = request.Title
	PKMBADME.Content = request.Content

	if err := c.PKMBADMERepository.Update(tx, PKMBADME); err != nil {
		c.Log.WithError(err).Error("error updating pkm BADME")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm BADME")
		return nil, err
	}

	return converter.PKMBADMEToResponse(PKMBADME), nil
}

func (c *PKMBADMEUseCase) Delete(ctx context.Context, request *model.DeletePKMBADMERequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	pkmBADME := new(entity.PKMBADME)
	if err := c.PKMBADMERepository.FindById(tx, pkmBADME, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkmBADME")
		return err
	}

	if err := c.PKMBADMERepository.Delete(tx, pkmBADME); err != nil {
		c.Log.WithError(err).Error("error deleting pkmBADME")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting pkmBADME")
		return err
	}

	return nil
}
