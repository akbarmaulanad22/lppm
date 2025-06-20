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

type PKMSTPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMSTPRepository *repository.PKMSTPRepository
}

func NewPKMSTPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMSTPRepository *repository.PKMSTPRepository,
) *PKMSTPUseCase {
	return &PKMSTPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMSTPRepository: PKMSTPRepository,
	}
}

func (c *PKMSTPUseCase) Create(ctx context.Context, request *model.CreatePKMSTPRequest) (*model.PKMSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMSTP := &entity.PKMSTP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMSTPRepository.Create(tx, PKMSTP); err != nil {
		c.Log.WithError(err).Error("failed to create pkm STP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMSTPToResponse(PKMSTP), nil
}

func (c *PKMSTPUseCase) FindAll(ctx context.Context) ([]model.PKMSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMSTP, err := c.PKMSTPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting pkm STP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting pkm STP")
		return nil, err
	}

	responses := make([]model.PKMSTPResponse, len(PKMSTP))
	for i, pkmStp := range PKMSTP {
		responses[i] = *converter.PKMSTPToResponse(&pkmStp)
	}

	return responses, nil
}

func (c *PKMSTPUseCase) Update(ctx context.Context, request *model.UpdatePKMSTPRequest) (*model.PKMSTPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMSTP := new(entity.PKMSTP)
	if err := c.PKMSTPRepository.FindById(tx, PKMSTP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkm STP")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMSTP.Title = request.Title
	PKMSTP.Content = request.Content

	if err := c.PKMSTPRepository.Update(tx, PKMSTP); err != nil {
		c.Log.WithError(err).Error("error updating pkm STP")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating pkm STP")
		return nil, err
	}

	return converter.PKMSTPToResponse(PKMSTP), nil
}

func (c *PKMSTPUseCase) Delete(ctx context.Context, request *model.DeletePKMSTPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	pkmSTP := new(entity.PKMSTP)
	if err := c.PKMSTPRepository.FindById(tx, pkmSTP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting pkmSTP")
		return err
	}

	if err := c.PKMSTPRepository.Delete(tx, pkmSTP); err != nil {
		c.Log.WithError(err).Error("error deleting pkmSTP")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting pkmSTP")
		return err
	}

	return nil
}
