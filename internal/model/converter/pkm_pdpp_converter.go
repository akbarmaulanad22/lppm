package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PKMPDPPToResponse(e *entity.PKMPDPP) *model.PKMPDPPResponse {
	return &model.PKMPDPPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
