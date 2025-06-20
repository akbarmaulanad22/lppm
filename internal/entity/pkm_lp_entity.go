package entity

type PKMLP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMLP) TableName() string {
	return "lppm_pkm_lp"
}
