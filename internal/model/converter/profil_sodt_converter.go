package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func ProfilSODTToResponse(e *entity.ProfilSODT) *model.ProfilSODTResponse {
	return &model.ProfilSODTResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
