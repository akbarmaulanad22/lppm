package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalTAJBToResponse(e *entity.JurnalTAJB) *model.JurnalTAJBResponse {
	return &model.JurnalTAJBResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
}
