package entity

type PenelitianSKR struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianSKR) TableName() string {
	return "lppm_penelitian_skr"
}
