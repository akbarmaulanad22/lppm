package entity

// ... existing code ...
type JurnalTeknois struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*JurnalTeknois) TableName() string {
	return "lppm_jurnal_teknois"
}
