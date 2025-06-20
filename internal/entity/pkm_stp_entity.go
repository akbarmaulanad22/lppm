package entity

type PKMSTP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMSTP) TableName() string {
	return "lppm_pkm_stp"
}
