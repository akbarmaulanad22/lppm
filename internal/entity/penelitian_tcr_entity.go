package entity

type PenelitianTCR struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianTCR) TableName() string {
	return "lppm_penelitian_tcr"
}
