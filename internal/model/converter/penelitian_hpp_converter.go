package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianHPPToResponse(e *entity.PenelitianHPP) *model.PenelitianHPPResponse {
	return &model.PenelitianHPPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 