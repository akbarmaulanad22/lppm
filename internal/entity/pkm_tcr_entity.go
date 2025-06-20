package entity

type PKMTCR struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMTCR) TableName() string {
	return "lppm_pkm_tcr"
}
