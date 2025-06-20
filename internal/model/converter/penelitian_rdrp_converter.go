package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianRDRPToResponse(e *entity.PenelitianRDRP) *model.PenelitianRDRPResponse {
	return &model.PenelitianRDRPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
