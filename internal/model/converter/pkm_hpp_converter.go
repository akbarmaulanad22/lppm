package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PKMHPPToResponse(e *entity.PKMHPP) *model.PKMHPPResponse {
	return &model.PKMHPPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
