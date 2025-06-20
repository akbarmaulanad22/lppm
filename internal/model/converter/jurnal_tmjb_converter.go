package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalTMJBToResponse(e *entity.JurnalTMJB) *model.JurnalTMJBResponse {
	return &model.JurnalTMJBResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 