package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianBADMEToResponse(e *entity.PenelitianBADME) *model.PenelitianBADMEResponse {
	return &model.PenelitianBADMEResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 