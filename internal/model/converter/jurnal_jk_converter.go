package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalJKToResponse(e *entity.JurnalJK) *model.JurnalJKResponse {
	return &model.JurnalJKResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
}
