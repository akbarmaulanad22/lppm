package converter

import (
	"tugasakhir/internal/entity"
	"tugasakhir/internal/model"
)

func JurnalJSToResponse(e *entity.JurnalJS) *model.JurnalJSResponse {
	return &model.JurnalJSResponse{
		ID:      e.ID,
		Title:   e.Title,
		Content: e.Content,
	}
} 