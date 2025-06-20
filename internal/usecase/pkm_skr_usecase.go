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

type PKMSKRUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMSKRRepository *repository.PKMSKRRepository
}

func NewPKMSKRUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMSKRRepository *repository.PKMSKRRepository,
) *PKMSKRUseCase {
	return &PKMSKRUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMSKRRepository: PKMSKRRepository,
	}
}

func (c *PKMSKRUseCase) Create(ctx context.Context, request *model.CreatePKMSKRRequest) (*model.PKMSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMSKR := &entity.PKMSKR{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMSKRRepository.Create(tx, PKMSKR); err != nil {
		c.Log.WithError(err).Error("failed to create pkm skr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMSKRToResponse(PKMSKR), nil
}

func (c *PKMSKRUseCase) FindAll(ctx context.Context) ([]model.PKMSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMSKR, err := c.PKMSKRRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm skr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm skr")
		return nil, err
	}

	responses := make([]model.PKMSKRResponse, len(PKMSKR))
	for i, pkmSkr := range PKMSKR {
		responses[i] = *converter.PKMSKRToResponse(&pkmSkr)
	}

	return responses, nil
}

func (c *PKMSKRUseCase) Update(ctx context.Context, request *model.UpdatePKMSKRRequest) (*model.PKMSKRResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMSKR := new(entity.PKMSKR)
	if err := c.PKMSKRRepository.FindById(tx, PKMSKR, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm skr")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMSKR.Title = request.Title
	PKMSKR.Content = request.Content

	if err := c.PKMSKRRepository.Update(tx, PKMSKR); err != nil {
		c.Log.WithError(err).Error("error updating pkm skr")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm skr")
		return nil, err
	}

	return converter.PKMSKRToResponse(PKMSKR), nil
}

func (c *PKMSKRUseCase) Delete(ctx context.Context, request *model.DeletePKMSKRRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	pkmSKR := new(entity.PKMSKR)
	if err := c.PKMSKRRepository.FindById(tx, pkmSKR, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkmSKR")
		return err
	}

	if err := c.PKMSKRRepository.Delete(tx, pkmSKR); err != nil {
		c.Log.WithError(err).Error("error deleting pkmSKR")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting pkmSKR")
		return err
	}

	return nil
}
