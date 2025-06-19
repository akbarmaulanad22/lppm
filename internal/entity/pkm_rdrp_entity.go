package entity

type PKMRDRP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMRDRP) TableName() string {
	return "lppm_pkm_rdrp"
}
