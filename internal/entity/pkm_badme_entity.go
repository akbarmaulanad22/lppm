package entity

type PKMBADME struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMBADME) TableName() string {
	return "lppm_pkm_badme"
}
