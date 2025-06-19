package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func ProfilVisiMisiToResponse(e *entity.ProfilVisiMisi) *model.ProfilVisiMisiResponse {
	return &model.ProfilVisiMisiResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}
