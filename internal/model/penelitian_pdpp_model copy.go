package model

type PenelitianPDPPResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type CreatePenelitianPDPPRequest struct {
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type UpdatePenelitianPDPPRequest struct {
	ID      uint   `json:"-"`
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type DeletePenelitianPDPPRequest struct {
	ID uint `json:"-" validate:"required"`
}
