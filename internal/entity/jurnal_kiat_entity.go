package entity

type JurnalKIAT struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*JurnalKIAT) TableName() string {
	return "lppm_jurnal_kiat"
}
