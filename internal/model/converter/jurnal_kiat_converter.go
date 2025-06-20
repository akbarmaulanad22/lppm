package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalKIATToResponse(e *entity.JurnalKIAT) *model.JurnalKIATResponse {
	return &model.JurnalKIATResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 