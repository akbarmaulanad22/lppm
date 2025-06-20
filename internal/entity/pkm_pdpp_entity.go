package entity

type PKMPDPP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMPDPP) TableName() string {
	return "lppm_pkm_pdpp"
}
