package entity

type JurnalTAJB struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*JurnalTAJB) TableName() string {
	return "lppm_jurnal_tajb"
}
