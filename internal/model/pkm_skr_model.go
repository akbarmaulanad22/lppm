package model

type PKMSKRResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type CreatePKMSKRRequest struct {
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type UpdatePKMSKRRequest struct {
	ID      uint   `json:"-"`
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type DeletePKMSKRRequest struct {
	ID uint `json:"-" validate:"required"`
}
