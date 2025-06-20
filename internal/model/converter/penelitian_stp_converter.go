package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func PenelitianSTPToResponse(e *entity.PenelitianSTP) *model.PenelitianSTPResponse {
	return &model.PenelitianSTPResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
}
