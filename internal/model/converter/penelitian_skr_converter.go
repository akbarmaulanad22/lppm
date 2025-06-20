package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianSKRToResponse(e *entity.PenelitianSKR) *model.PenelitianSKRResponse {
	return &model.PenelitianSKRResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 