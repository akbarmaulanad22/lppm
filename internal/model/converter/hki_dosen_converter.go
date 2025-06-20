package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func HKIDosenToResponse(e *entity.HKIDosen) *model.HKIDosenResponse {
	return &model.HKIDosenResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
