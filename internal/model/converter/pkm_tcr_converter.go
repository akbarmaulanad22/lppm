package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PKMTCRToResponse(e *entity.PKMTCR) *model.PKMTCRResponse {
	return &model.PKMTCRResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
