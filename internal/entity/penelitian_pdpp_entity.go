package entity

type PenelitianPDPP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianPDPP) TableName() string {
	return "lppm_penelitian_pdpp"
}
