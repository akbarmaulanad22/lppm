package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalTeknoisToResponse(e *entity.JurnalTeknois) *model.JurnalTeknoisResponse {
	return &model.JurnalTeknoisResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 