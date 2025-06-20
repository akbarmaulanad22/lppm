package entity

type PKMHPP struct {
	Entity
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (*PKMHPP) TableName() string {
	return "lppm_pkm_hpp"
}
