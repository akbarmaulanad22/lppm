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

type PKMLPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMLPRepository *repository.PKMLPRepository
}

func NewPKMLPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMLPRepository *repository.PKMLPRepository,
) *PKMLPUseCase {
	return &PKMLPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMLPRepository: PKMLPRepository,
	}
}

func (c *PKMLPUseCase) Create(ctx context.Context, request *model.CreatePKMLPRequest) (*model.PKMLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMLP := &entity.PKMLP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMLPRepository.Create(tx, PKMLP); err != nil {
		c.Log.WithError(err).Error("failed to create pkm LP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMLPToResponse(PKMLP), nil
}

func (c *PKMLPUseCase) FindAll(ctx context.Context) ([]model.PKMLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMLP, err := c.PKMLPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm LP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm LP")
		return nil, err
	}

	responses := make([]model.PKMLPResponse, len(PKMLP))
	for i, pkmLp := range PKMLP {
		responses[i] = *converter.PKMLPToResponse(&pkmLp)
	}

	return responses, nil
}

func (c *PKMLPUseCase) Update(ctx context.Context, request *model.UpdatePKMLPRequest) (*model.PKMLPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMLP := new(entity.PKMLP)
	if err := c.PKMLPRepository.FindById(tx, PKMLP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm LP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMLP.Title = request.Title
	PKMLP.Content = request.Content

	if err := c.PKMLPRepository.Update(tx, PKMLP); err != nil {
		c.Log.WithError(err).Error("error updating pkm LP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm LP")
		return nil, err
	}

	return converter.PKMLPToResponse(PKMLP), nil
}

func (c *PKMLPUseCase) Delete(ctx context.Context, request *model.DeletePKMLPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	pkmLP := new(entity.PKMLP)
	if err := c.PKMLPRepository.FindById(tx, pkmLP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkmLP")
		return err
	}

	if err := c.PKMLPRepository.Delete(tx, pkmLP); err != nil {
		c.Log.WithError(err).Error("error deleting pkmLP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting pkmLP")
		return err
	}

	return nil
}
