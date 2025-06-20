package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PKMLPToResponse(e *entity.PKMLP) *model.PKMLPResponse {
	return &model.PKMLPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
