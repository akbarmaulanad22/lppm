package entity

type PenelitianLP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianLP) TableName() string {
	return "lppm_penelitian_lp"
} 