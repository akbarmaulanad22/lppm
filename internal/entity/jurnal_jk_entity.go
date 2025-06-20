package entity

type JurnalJK struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*JurnalJK) TableName() string {
	return "lppm_jurnal_jk"
}
