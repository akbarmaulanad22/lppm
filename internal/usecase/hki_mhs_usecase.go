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

type HKIMHSUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	HKIMHSRepository *repository.HKIMHSRepository
}

func NewHKIMHSUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	HKIMHSRepository *repository.HKIMHSRepository,
) *HKIMHSUseCase {
	return &HKIMHSUseCase{
		DB:               db,
		Log:              logger,
		Validate:         validate,
		HKIMHSRepository: HKIMHSRepository,
	}
}

func (c *HKIMHSUseCase) Create(ctx context.Context, request *model.CreateHKIMHSRequest) (*model.HKIMHSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	HKIMHS := &entity.HKIMHS{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.HKIMHSRepository.Create(tx, HKIMHS); err != nil {
		c.Log.WithError(err).Error("failed to create HKI Mahasiswa")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.HKIMHSToResponse(HKIMHS), nil
}

func (c *HKIMHSUseCase) FindAll(ctx context.Context) ([]model.HKIMHSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	HKIMHS, err := c.HKIMHSRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting HKI Mahasiswa")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting HKI Mahasiswa")
		return nil, err
	}

	responses := make([]model.HKIMHSResponse, len(HKIMHS))
	for i, hkiMhs := range HKIMHS {
		responses[i] = *converter.HKIMHSToResponse(&hkiMhs)
	}

	return responses, nil
}

func (c *HKIMHSUseCase) Update(ctx context.Context, request *model.UpdateHKIMHSRequest) (*model.HKIMHSResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	HKIMHS := new(entity.HKIMHS)
	if err := c.HKIMHSRepository.FindById(tx, HKIMHS, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting HKI Mahasiswa")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	HKIMHS.Title = request.Title
	HKIMHS.Content = request.Content

	if err := c.HKIMHSRepository.Update(tx, HKIMHS); err != nil {
		c.Log.WithError(err).Error("error updating HKI Mahasiswa")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating HKI Mahasiswa")
		return nil, err
	}

	return converter.HKIMHSToResponse(HKIMHS), nil
}

func (c *HKIMHSUseCase) Delete(ctx context.Context, request *model.DeleteHKIMHSRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	hkiMHS := new(entity.HKIMHS)
	if err := c.HKIMHSRepository.FindById(tx, hkiMHS, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting hkiMHS")
		return err
	}

	if err := c.HKIMHSRepository.Delete(tx, hkiMHS); err != nil {
		c.Log.WithError(err).Error("error deleting hkiMHS")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting hkiMHS")
		return err
	}

	return nil
}
