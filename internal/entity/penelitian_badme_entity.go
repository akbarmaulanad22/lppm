package entity

type PenelitianBADME struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianBADME) TableName() string {
	return "lppm_penelitian_badme"
} 