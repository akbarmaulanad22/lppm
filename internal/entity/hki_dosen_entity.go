package entity

type HKIDosen struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*HKIDosen) TableName() string {
	return "lppm_hki_dosen"
}
