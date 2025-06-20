package entity

type PKMSKR struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMSKR) TableName() string {
	return "lppm_pkm_skr"
}
