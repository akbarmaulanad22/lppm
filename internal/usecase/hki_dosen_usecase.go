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

type HKIDosenUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	HKIDosenRepository *repository.HKIDosenRepository
}

func NewHKIDosenUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	HKIDosenRepository *repository.HKIDosenRepository,
) *HKIDosenUseCase {
	return &HKIDosenUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		HKIDosenRepository: HKIDosenRepository,
	}
}

func (c *HKIDosenUseCase) Create(ctx context.Context, request *model.CreateHKIDosenRequest) (*model.HKIDosenResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	HKIDosen := &entity.HKIDosen{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := c.HKIDosenRepository.Create(tx, HKIDosen); err != nil {
		c.Log.WithError(err).Error("failed to create HKI Dosen")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.HKIDosenToResponse(HKIDosen), nil
}

func (c *HKIDosenUseCase) FindAll(ctx context.Context) ([]model.HKIDosenResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	HKIDosen, err := c.HKIDosenRepository.FindAll(tx)
	if err != nil {
		c.Log.WithError(err).Error("error getting HKI Dosen")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting HKI Dosen")
		return nil, err
	}

	responses := make([]model.HKIDosenResponse, len(HKIDosen))
	for i, hkiDosen := range HKIDosen {
		responses[i] = *converter.HKIDosenToResponse(&hkiDosen)
	}

	return responses, nil
}

func (c *HKIDosenUseCase) Update(ctx context.Context, request *model.UpdateHKIDosenRequest) (*model.HKIDosenResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	HKIDosen := new(entity.HKIDosen)
	if err := c.HKIDosenRepository.FindById(tx, HKIDosen, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting HKI Dosen")
		return nil, err
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, err
	}

	HKIDosen.Title = request.Title
	HKIDosen.Content = request.Content

	if err := c.HKIDosenRepository.Update(tx, HKIDosen); err != nil {
		c.Log.WithError(err).Error("error updating HKI Dosen")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating HKI Dosen")
		return nil, err
	}

	return converter.HKIDosenToResponse(HKIDosen), nil
}

func (c *HKIDosenUseCase) Delete(ctx context.Context, request *model.DeleteHKIDosenRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return err
	}

	hkiDOSEN := new(entity.HKIDosen)
	if err := c.HKIDosenRepository.FindById(tx, hkiDOSEN, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting hkiDOSEN")
		return err
	}

	if err := c.HKIDosenRepository.Delete(tx, hkiDOSEN); err != nil {
		c.Log.WithError(err).Error("error deleting hkiDOSEN")
		return err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting hkiDOSEN")
		return err
	}

	return nil
}
