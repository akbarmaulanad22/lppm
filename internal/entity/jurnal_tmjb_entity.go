package entity

type JurnalTMJB struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*JurnalTMJB) TableName() string {
	return "lppm_jurnal_tmjb"
}
