package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianPDPPToResponse(e *entity.PenelitianPDPP) *model.PenelitianPDPPResponse {
	return &model.PenelitianPDPPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
