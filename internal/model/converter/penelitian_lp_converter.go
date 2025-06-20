package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianLPToResponse(e *entity.PenelitianLP) *model.PenelitianLPResponse {
	return &model.PenelitianLPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
}
