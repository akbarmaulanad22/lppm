package entity

type HKIMHS struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*HKIMHS) TableName() string {
	return "lppm_hki_mhs"
}
