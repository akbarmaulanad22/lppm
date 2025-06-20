package entity

type PenelitianSTP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PenelitianSTP) TableName() string {
	return "lppm_penelitian_stp"
}
