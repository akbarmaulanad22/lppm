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

type PKMRDRPUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PKMRDRPRepository *repository.PKMRDRPRepository
}

func NewPKMRDRPUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PKMRDRPRepository *repository.PKMRDRPRepository,
) *PKMRDRPUseCase {
	return &PKMRDRPUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PKMRDRPRepository: PKMRDRPRepository,
	}
}

func (c *PKMRDRPUseCase) Create(ctx context.Context, request *model.CreatePKMRDRPRequest) (*model.PKMRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	PKMRDRP := &entity.PKMRDRP{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.PKMRDRPRepository.Create(tx, PKMRDRP); err != nil {
		c.Log.WithError(err).Error("failed to create profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.PKMRDRPToResponse(PKMRDRP), nil
}

func (c *PKMRDRPUseCase) FindAll(ctx context.Context) ([]model.PKMRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMRDRP, err := c.PKMRDRPRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	responses := make([]model.PKMRDRPResponse, len(PKMRDRP))
	for i, visiMisi := range PKMRDRP {
		responses[i] = *converter.PKMRDRPToResponse(&visiMisi)
	}

	return responses, nil
}

func (c *PKMRDRPUseCase) Update(ctx context.Context, request *model.UpdatePKMRDRPRequest) (*model.PKMRDRPResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	PKMRDRP := new(entity.PKMRDRP)
	if err := c.PKMRDRPRepository.FindById(tx, PKMRDRP, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting profil Visi Misi")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	PKMRDRP.Title = request.Title
	PKMRDRP.Content = request.Content

	if err := c.PKMRDRPRepository.Update(tx, PKMRDRP); err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating profil Visi Misi")
		return nil, err
	}

	return converter.PKMRDRPToResponse(PKMRDRP), nil
}

func (c *PKMRDRPUseCase) Delete(ctx context.Context, request *model.DeletePKMRDRPRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	contact := new(entity.PKMRDRP)
	if err := c.PKMRDRPRepository.FindById(tx, contact, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	if err := c.PKMRDRPRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return nil
}
