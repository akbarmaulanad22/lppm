package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func HKIMHSToResponse(e *entity.HKIMHS) *model.HKIMHSResponse {
	return &model.HKIMHSResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
