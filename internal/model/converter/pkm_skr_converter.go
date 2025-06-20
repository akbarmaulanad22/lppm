package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PKMSKRToResponse(e *entity.PKMSKR) *model.PKMSKRResponse {
	return &model.PKMSKRResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
